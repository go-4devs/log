package log_test

import (
	"path/filepath"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/level"
)

func ExampleNew_withCaller() {
	logger := log.New(log.WithStdout()).With(
		log.WithLevel(log.KeyLevel, level.Debug),
		log.WithSource(3, filepath.Base),
	)
	logger.Err(ctx, "same error message")
	logger.InfoKVs(ctx, "same info message", "api-version", 0.1)
	_, _ = logger.Write([]byte("same write message"))

	// Output:
	// msg="same error message" level=error source=logger_example_caller_test.go:15
	// msg="same info message" api-version=0.1 level=info source=logger_example_caller_test.go:16
	// msg="same write message" level=info source=logger_example_caller_test.go:17
}
