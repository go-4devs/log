package log_test

import (
	"io"

	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/handler/zap"
	uzap "go.uber.org/zap"
)

func ExampleNew_zapHandler() {
	log := zap.New(zap.WithZap(uzap.NewExample()))
	log.Err(ctx, "log zap")
	log.ErrKV(ctx, "log zap kv", field.Int("int", 42))
	log.ErrKVs(ctx, "log zap kv sugar", "err", io.EOF)

	// Output:
	// {"level":"error","msg":"log zap"}
	// {"level":"error","msg":"log zap kv","int":42}
	// {"level":"error","msg":"log zap kv sugar","err":"EOF"}
}
