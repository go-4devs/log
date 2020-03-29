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
	// Output: level: info; msg: same message;
}

func ExampleNew_errf() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelInfo))
	logger.Errf(ctx, "same message %d", 1)
	// Output: level: error; msg: same message 1;
}

func ExampleNew_debugKV() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelDebug))
	logger.DebugKV(ctx, "same message", "error", os.ErrNotExist)
	// Output: level: debug; msg: same message; error: file does not exist;
}

func ExampleNew_level() {
	ctx := context.Background()

	logger := New(NewStdHandler(log.New(os.Stdout, "", 0), LevelError))
	logger.Info(ctx, "same message")
	// Output:

	logger.Err(ctx, "same error message")
	// Output: level: error; msg: same error message;
}

type ctxKey string

func ctxProcessor(key ctxKey) func(ctx context.Context) Field {
	return func(ctx context.Context) Field {
		return Field{Key: string(key), Value: ctx.Value(key)}
	}
}
func goVersion(ctx context.Context) Field {
	return Field{Key: "go", Value: "go1.13.4"}
}

func apiVersion(ctx context.Context) Field {
	return Field{Key: "api", Value: "0.1.0"}
}

func ExampleProcessor() {
	var ctxKey ctxKey = "requestID"
	ctx := context.WithValue(context.Background(), ctxKey, "6a5fa048-7181-11ea-bc55-0242ac130003")
	l := log.New(os.Stdout, "", 0)

	logger := New(
		NewStdHandler(l, LevelInfo),
		WithProcessor(ctxProcessor(ctxKey), apiVersion),
		WithProcessor(goVersion),
	)
	logger.Info(ctx, "same message")
	// Output: level: info; msg: same message; requestID: 6a5fa048-7181-11ea-bc55-0242ac130003; api: 0.1.0; go: go1.13.4;
}
