package log_test

import (
	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/level"
)

func ExampleNew_withCaller() {
	logger := log.With(
		log.New(log.WithStdout()),
		log.WithLevel("level", level.Debug),
		log.WithCaller("caller", 2, false),
	)
	logger.Err(ctx, "same error message")
	logger.InfoKVs(ctx, "same info message", "api-version", 0.1)
	_, _ = logger.Write([]byte("same write message"))

	// Output:
	// msg="same error message" level=error caller=logger_example_caller_test.go:14
	// msg="same info message" api-version=0.1 level=info caller=logger_example_caller_test.go:15
	// msg="same write message" level=info caller=logger_example_caller_test.go:16
}
