package log

import (
	"context"
	"log"
	"os"
)

func ExampleNew() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelInfo))
	logger.Info(ctx, "same message")
	// Output: msg="same message"
}

func ExampleNew_errf() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelInfo))
	logger.Errf(ctx, "same message %d", 1)
	// Output: msg="same message 1"
}

func ExampleNew_debugKV() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelDebug), levelInfo)
	logger.DebugKV(ctx, "same message", "error", os.ErrNotExist)
	// Output: msg="same message" error=file does not exist level=debug
}

func ExampleNew_level() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelError), levelInfo)
	logger.Info(ctx, "same message")
	// Output:

	logger.Err(ctx, "same error message")
	// Output: msg="same error message" level=error
}

type ctxKey string

func ctxProcessor(key ctxKey) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Handler) {
		handler(ctx, level, msg, append(fields, Field{Key: string(key), Value: ctx.Value(key)}))
	}
}

func goVersion(ctx context.Context, level Level, msg string, fields Fields, handler Handler) {
	handler(ctx, level, msg, append(fields, Field{Key: "go", Value: "go1.13.4"}))
}

func apiVersion(ctx context.Context, level Level, msg string, fields Fields, handler Handler) {
	handler(ctx, level, msg, append(fields, Field{Key: "api", Value: "0.1.0"}))
}

func levelInfo(ctx context.Context, level Level, msg string, fields Fields, handler Handler) {
	handler(ctx, level, msg, append(fields, Field{Key: "level", Value: level}))
}

func ExampleProcessor() {
	var ctxKey ctxKey = "requestID"
	ctx := context.WithValue(context.Background(), ctxKey, "6a5fa048-7181-11ea-bc55-0242ac130003")
	l := log.New(os.Stdout, "", 0)

	logger := New(
		NewStdHandler(l, LevelInfo),
		levelInfo, ctxProcessor(ctxKey), apiVersion, goVersion,
	)
	logger.Info(ctx, "same message")
	// Output: msg="same message" level=info requestID=6a5fa048-7181-11ea-bc55-0242ac130003 api=0.1.0 go=go1.13.4
}
