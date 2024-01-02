package log_test

import (
	"context"

	"gitoa.ru/go-4devs/log"
)

func ExampleWithSource() {
	ctx := context.Background()
	logger := log.New(log.WithStdout()).With(log.WithSource(1))

	logger.Debug(ctx, "debug message")
	// Output:
	// msg="debug message" source=source_example_test.go:13
}

func ExampleWithSource_json() {
	ctx := context.Background()
	logger := log.New(log.WithStdout(), log.WithJSONFormat()).With(log.WithSource(1))

	logger.Debug(ctx, "debug message")
	// Output:
	// {"msg":"debug message","source":{"file":"source_example_test.go","line":22,"func":"log_test.ExampleWithSource_json"}}
}
