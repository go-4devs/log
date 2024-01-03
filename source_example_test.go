package log_test

import (
	"context"
	"path/filepath"

	"gitoa.ru/go-4devs/log"
)

func ExampleWithSource() {
	ctx := context.Background()
	logger := log.New(log.WithStdout()).With(log.WithSource(1, filepath.Base))

	logger.Debug(ctx, "debug message")
	// Output:
	// msg="debug message" source=source_example_test.go:14
}

func ExampleWithSource_json() {
	ctx := context.Background()
	logger := log.New(log.WithStdout(), log.WithJSONFormat()).With(log.WithSource(2, filepath.Base))

	logger.Debug(ctx, "debug message")
	// Output:
	// {"msg":"debug message","source":{"file":"source_example_test.go","line":23,"func":"log_test.ExampleWithSource_json"}}
}
