package log_test

import (
	"bytes"
	"context"
	"os"
	"sync/atomic"
	"testing"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
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
		With(log.WithLevel("level", level.Info))
	success := "msg=message err=\"file already exists\" version=0.1.0 obj={id:uid} closure=\"some closure data\" level=info\n"

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

	success := "msg=\"info message\" err=\"file already exists\" requestID=6a5fa048-7181-11ea-bc55-0242ac1311113 level=info\n"
	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf)).With(log.WithContextValue(requestID), log.WithLevel("level", level.Info))

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

func TestLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf)).With(log.WithContextValue(requestID), log.WithLevel("level", level.Info))

	_, err := logger(ctx, nil)
	if err != nil {
		t.Fatalf("expected <nil> err, got: %v", err)
	}

	if buf.String() != "" {
		t.Errorf("invalid value\n got:%+v\n exp:\"\"", buf)
	}

	_, err = logger(ctx, entry.New().SetLevel(level.Error))
	if err != nil {
		t.Fatalf("expected <nil> err, got: %v", err)
	}

	success := "msg=\"\" requestID=<nil> level=error\n"

	if buf.String() != success {
		t.Errorf("invalid value\n got:%+v\n exp:%+v", buf, success)
	}
}
