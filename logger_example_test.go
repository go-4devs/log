package log_test

import (
	"context"
	"fmt"
	std "log"
	"os"

	"gitoa.ru/go-4devs/log"
)

//nolint:gochecknoglobals
var ctx = context.Background()

func ExampleNew() {
	logger := log.New(log.WithStdout())
	logger.Info(ctx, "same message")
	// Output: msg="same message"
}

func ExampleInfo() {
	std.SetOutput(os.Stdout)
	std.SetFlags(0)
	log.Info(ctx, "same message")
	// Output: msg="same message" level=info
}

func ExampleErrKV() {
	std.SetOutput(os.Stdout)
	std.SetFlags(0)
	log.ErrKV(ctx, "same message", "key", "addition value")
	// Output: msg="same message" key=addition value level=error
}

func ExampleNew_errf() {
	logger := log.New(log.WithStdout())
	logger.Errf(ctx, "same message %d", 1)
	// Output: msg="same message 1"
}

func ExampleNew_debugKV() {
	logger := log.New(log.WithStdout()).With(log.WithLevel(log.LevelDebug))
	logger.DebugKV(ctx, "same message", "error", os.ErrNotExist)
	// Output: msg="same message" error=file does not exist level=debug
}

func ExampleNew_level() {
	logger := log.New(log.WithStdout()).With(log.WithLevel(log.LevelError))
	logger.Info(ctx, "same message")
	// Output:

	logger.Err(ctx, "same error message")
	// Output: msg="same error message" level=error
}

func ExampleNew_jsonFormat() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat()).
		With(
			log.WithCaller(4, true),
			log.WithLevel(log.LevelDebug),
			log.GoVersion("go-version"),
		)
	logger.Err(ctx, "same error message")
	// Output: {"caller":"logger_example_test.go:63","go-version":"go1.14.2","level":"error","msg":"same error message"}
}

func ExampleNew_withLogger() {
	stdlogger := std.New(os.Stdout, "same prefix ", std.Lshortfile)
	logger := log.With(
		log.New(
			log.WithLogger(
				stdlogger,
				func(msg string, fields log.Fields) string {
					return fmt.Sprint("msg=\"", msg, "\" ", fields)
				},
			),
			log.WithCalldepth(9),
		),
		log.WithLevel(log.LevelDebug),
		log.GoVersion("go-version"),
	)
	logger.Err(ctx, "same error message")
	logger.InfoKV(ctx, "same info message", "api-version", 0.1)

	// Output:
	// same prefix logger_example_test.go:82: msg="same error message" level=error go-version=go1.14.2
	// same prefix logger_example_test.go:83: msg="same info message" api-version=0.1 level=info go-version=go1.14.2
}

type ctxKey string

func (c ctxKey) String() string {
	return string(c)
}

func levelInfo(ctx context.Context, level log.Level, msg string, fields log.Fields, handler log.Logger) {
	handler(ctx, level, msg, append(fields, log.Field{Key: "level", Value: level}))
}

func ExampleWith() {
	var requestID ctxKey = "requestID"
	vctx := context.WithValue(ctx, requestID, "6a5fa048-7181-11ea-bc55-0242ac130003")

	logger := log.With(
		log.New(log.WithStdout()),
		levelInfo, log.WithContextValue(requestID), log.KeyValue("api", "0.1.0"), log.GoVersion("go"),
	)
	logger.Info(vctx, "same message")
	// Output: msg="same message" level=info requestID=6a5fa048-7181-11ea-bc55-0242ac130003 api=0.1.0 go=go1.14.2
}

func ExampleLogger_Print() {
	logger := log.With(
		log.New(log.WithStdout()),
		levelInfo, log.KeyValue("client", "http"), log.KeyValue("api", "0.1.0"), log.GoVersion("go"),
	)
	logger.Print("same message")
	// Output: msg="same message" level=info client=http api=0.1.0 go=go1.14.2
}

func ExamplePrint() {
	std.SetOutput(os.Stdout)
	std.SetFlags(0)
	log.Print("same message")
	// Output: msg="same message" level=info
}
