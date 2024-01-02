package log_test

import (
	"context"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/level"
)

func ExampleDebug() {
	logger := log.With(log.New(log.WithStdout()),
		log.WithSource(2),
		log.WithLevel(log.KeyLevel, level.Debug),
		log.WithExit(level.Alert),
		log.WithPanic(level.Emergency),
	)

	log.SetLogger(logger)

	ctx := context.Background()
	log.Debug(ctx, "debug message")
	// Output:
	// msg="debug message" source=global_example_test.go:21 level=debug
}
