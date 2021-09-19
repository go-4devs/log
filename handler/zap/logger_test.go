package zap_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	zlog "gitoa.ru/go-4devs/log/handler/zap"
	"gitoa.ru/go-4devs/log/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
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
	logger := zlog.New(zlog.WithZap(zap.New(core)))
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
