package zap_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	zlog "github.com/go-4devs/log/handler/zap"
	"github.com/go-4devs/log/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := &bytes.Buffer{}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}), zapcore.AddSync(buf), zapcore.DebugLevel)
	logger := zlog.New(zap.New(core))
	expect := `{"level":"info","msg":"handle zap message","env":"test"}` + "\n"

	if _, err := logger(ctx, entry.New(
		entry.WithFields(field.String("env", "test")),
		entry.WithLevel(level.Notice),
		entry.WithMessage("handle zap message"),
	)); err != nil {
		t.Error(err)
	}

	if buf.String() != expect {
		t.Errorf("invalid message\n got: %s\nexpect:%s\n", buf.String(), expect)
	}
}

func TestNewSevereLevelsDoNotStop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := &bytes.Buffer{}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}), zapcore.AddSync(buf), zapcore.DebugLevel)
	logger := zlog.New(zap.New(core))

	cases := []struct {
		lvl  level.Level
		want string
	}{
		{lvl: level.Emergency, want: "panic"},
		{lvl: level.Alert, want: "fatal"},
	}

	for _, tc := range cases {
		buf.Reset()

		_, err := logger(ctx, entry.New(
			entry.WithLevel(tc.lvl),
			entry.WithMessage("stop"),
		))
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `{"level":"` + tc.want + `","msg":"stop"}` + "\n"
		if got != want {
			t.Fatalf("%s: got %s", tc.lvl, got)
		}
	}
}
