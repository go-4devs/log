package log_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
)

var requestID ctxKey = "requestID"

func TestFields(t *testing.T) {
	t.Parallel()

	type rObj struct {
		id string
	}

	var cnt int32

	ctx := context.Background()
	buf := &bytes.Buffer{}
	log := log.New(log.WithWriter(buf)).
		With(log.WithLevel(log.KeyLevel, level.Info))
	success := "msg=message err=\"file already exists\"" +
		" version=0.1.0 obj={id:uid} closure=\"some closure data\" level=info\n"

	log.InfoKVs(ctx, "message",
		"err", os.ErrExist,
		"version", "0.1.0",
		"obj", rObj{id: "uid"},
		"closure", func() any {
			atomic.AddInt32(&cnt, 1)

			return "some closure data"
		},
	)

	log.DebugKVs(ctx, "debug message",
		"closure", field.ClosureFn(func() any {
			atomic.AddInt32(&cnt, 1)

			return "some debug data"
		}),
	)

	if success != buf.String() {
		t.Errorf("invalid value\n got:%s\n exp:%s", buf, success)
	}

	if cnt != 1 {
		t.Errorf("invalid cnt value\n got:%d\n exp:1", cnt)
	}
}

func TestWriter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	success := "msg=\"info message\" err=\"file already exists\"" +
		" requestID=6a5fa048-7181-11ea-bc55-0242ac1311113 level=info\n"
	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf)).With(log.WithContextValue(requestID), log.WithLevel(log.KeyLevel, level.Info))

	_, _ = logger.Writer(
		context.WithValue(ctx, requestID, "6a5fa048-7181-11ea-bc55-0242ac1311113"),
		level.Info,
		field.Error("err", os.ErrExist),
	).Write([]byte("info message"))

	if success != buf.String() {
		t.Errorf("invalid value\n got:%s\n exp:%s", buf, success)
	}

	buf.Reset()

	_, _ = logger.Writer(ctx, level.Debug).Write([]byte("debug message"))

	if buf.String() != "" {
		t.Errorf("invalid value\n got:%s\n exp:%s", buf, success)
	}
}

func TestNewConcurrentWrite(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf))

	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 40 {
				logger.Info(context.Background(), "x")
			}
		})
	}

	wg.Wait()
}

func TestWithMetrics(t *testing.T) {
	t.Parallel()

	var got level.Level

	logger := log.With(
		func(_ context.Context, e *entry.Entry) (int, error) {
			if got != level.Error {
				t.Errorf("metrics saw %s before handler, want error", got)
			}

			e.SetLevel(level.Debug)

			return 0, nil
		},
		log.WithMetrics(func(lvl level.Level) {
			got = lvl
		}),
	)

	_, err := logger(context.Background(), entry.New().SetLevel(level.Error))
	if err != nil {
		t.Fatal(err)
	}

	if got != level.Error {
		t.Errorf("got %s, want error", got)
	}
}

func TestLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf)).With(log.WithContextValue(requestID), log.WithLevel(log.KeyLevel, level.Info))

	_, err := logger(ctx, entry.New().SetLevel(level.Error))
	if err != nil {
		t.Fatalf("expected <nil> err, got: %v", err)
	}

	success := "msg=\"\" requestID=<nil> level=error\n"

	if buf.String() != success {
		t.Errorf("invalid value\n got:%+v\n exp:%+v", buf, success)
	}
}

func TestHandlerKeepsEntry(t *testing.T) {
	t.Parallel()

	var first *entry.Entry

	var calls int

	logger := log.Logger(func(_ context.Context, item *entry.Entry) (int, error) {
		if calls == 0 {
			first = item
		}

		calls++

		return 0, nil
	})

	logger.Info(context.Background(), "one")
	logger.Info(context.Background(), "two")

	if first.Message() != "one" {
		t.Fatalf("kept %q", first.Message())
	}
}

func TestPaniclnPanicsTrimmedMessage(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf))

	defer func() {
		if recover() != "hello there" || buf.String() != "msg=\"hello there\"\n" {
			t.Fatalf("panic log %q", buf.String())
		}
	}()

	logger.Panicln("hello", "there")
	t.Fatal("did not panic")
}

func TestPrintlnHasNoTrailingNewline(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	log.New(log.WithWriter(buf)).Println("hello", "there")

	if buf.String() != "msg=\"hello there\"\n" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestTwoLoggersShareWriter(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	left := log.New(log.WithWriter(buf))
	right := log.New(log.WithWriter(buf))

	var wg sync.WaitGroup

	const lines = 40

	wg.Add(lines * 2)

	for range lines {
		go func() {
			defer wg.Done()

			left.Info(context.Background(), "aaaaaaaa")
		}()
		go func() {
			defer wg.Done()

			right.Info(context.Background(), "bbbbbbbb")
		}()
	}

	wg.Wait()

	for line := range bytes.SplitSeq(bytes.TrimSuffix(buf.Bytes(), []byte("\n")), []byte("\n")) {
		if bytes.Contains(line, []byte("aaaaaaaa")) && bytes.Contains(line, []byte("bbbbbbbb")) {
			t.Fatalf("torn %q", line)
		}
	}
}

//nolint:paralleltest // replaces the process logger
func TestGlobalSevereLevelsDoNotStop(t *testing.T) {
	previous := log.Log()
	buf := &bytes.Buffer{}

	log.SetLogger(log.New(log.WithWriter(buf)))
	t.Cleanup(func() { log.SetLogger(previous) })

	log.Emerg(context.Background(), "disk")
	log.Alert(context.Background(), "notice")
	log.SetLogger(nil)
	log.Info(context.Background(), "kept")

	got := buf.String()
	for _, part := range []string{"disk", "notice", "kept"} {
		if !strings.Contains(got, part) {
			t.Fatalf("%q missing %s", got, part)
		}
	}
}
