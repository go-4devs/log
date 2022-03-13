package log_test

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

//nolint:gochecknoglobals
var ctx = context.Background()

func setStdout() {
	// set stout for example by default stderror
	log.SetLogger(log.New(log.WithStdout()).With(log.WithLevel("level", level.Debug)))
}

func ExampleNew() {
	logger := log.New(log.WithStdout())
	logger.Info(ctx, "same message")
	// Output: msg="same message"
}

func ExampleInfo() {
	setStdout()
	log.Info(ctx, "same message")
	// Output: msg="same message" level=info
}

func ExampleErrKV() {
	setStdout()
	log.ErrKVs(ctx, "same message", "key", "addition value")
	// Output: msg="same message" key=addition value level=error
}

func ExampleNew_errf() {
	logger := log.New(log.WithStdout())
	logger.Errf(ctx, "same message %d", 1)
	// Output: msg="same message 1"
}

func ExampleNew_debugKV() {
	logger := log.New(log.WithStdout()).With(log.WithLevel("level", level.Debug))
	logger.DebugKVs(ctx, "same message", "error", os.ErrNotExist)
	// Output: msg="same message" error=file does not exist level=debug
}

func ExampleNew_level() {
	logger := log.New(log.WithStdout()).With(log.WithLevel("level", level.Error))
	logger.Err(ctx, "same error message")
	// Output: msg="same error message" level=error
}

func ExampleNew_level_info() {
	logger := log.New(log.WithStdout()).With(log.WithLevel("level", level.Error))
	logger.Info(ctx, "same message")
	// Output:
}

func ExampleNew_jsonFormat() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat()).
		With(
			log.WithLevel("level", level.Debug),
			log.GoVersion("go-version"),
		)
	logger.Err(ctx, "same error message")
	// Output: {"go-version":"go1.17.8","level":"error","msg":"same error message"}
}

func ExampleNew_textEncoding() {
	logger := log.With(
		log.New(log.WithStdout()),
		log.WithLevel("level", level.Debug),
		log.GoVersion("go-version"),
	)
	logger.Err(ctx, "same error message")
	logger.InfoKVs(ctx, "same info message", "api-version", 0.1)

	// Output:
	// msg="same error message" level=error go-version=go1.17.8
	// msg="same info message" api-version=0.1 level=info go-version=go1.17.8
}

type ctxKey string

func (c ctxKey) String() string {
	return string(c)
}

func levelInfo(ctx context.Context, entry *entry.Entry, handler log.Logger) (int, error) {
	return handler(ctx, entry.Add(field.String("level", entry.Level().String())))
}

func ExampleWith() {
	var requestID ctxKey = "requestID"
	vctx := context.WithValue(ctx, requestID, "6a5fa048-7181-11ea-bc55-0242ac130003")

	logger := log.With(
		log.New(log.WithStdout()),
		levelInfo, log.WithContextValue(requestID), log.KeyValue("api", "0.1.0"), log.GoVersion("go"),
	)
	logger.Info(vctx, "same message")
	// Output: msg="same message" level=info requestID=6a5fa048-7181-11ea-bc55-0242ac130003 api=0.1.0 go=go1.17.8
}

func ExampleLogger_Print() {
	logger := log.With(
		log.New(log.WithStdout()),
		levelInfo, log.KeyValue("client", "http"), log.KeyValue("api", "0.1.0"), log.GoVersion("go"),
	)
	logger.Print("same message")
	// Output: msg="same message" level=info client=http api=0.1.0 go=go1.17.8
}

func ExamplePrint() {
	setStdout()
	log.Print("same message")
	// Output: msg="same message" level=info
}

func ExampleWithClosure() {
	cnt := int32(0)
	closure := func() string {
		d := fmt.Sprintf("additional error data: %d", cnt)
		atomic.AddInt32(&cnt, 1)

		return d
	}

	log := log.With(log.New(log.WithStdout()), log.WithLevel("level", level.Info), log.WithClosure)

	log.DebugKVs(ctx, "debug message", "data", closure)
	log.ErrKVs(ctx, "error message", "err", closure)
	log.WarnKVs(ctx, "warn message", "warn", closure)

	// Output:
	// msg="error message" err=additional error data: 0 level=error
	// msg="warn message" warn=additional error data: 1 level=warning
}
