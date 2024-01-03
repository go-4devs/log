package log_test

import (
	"context"
	"math"
	"path/filepath"
	"time"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

func exampleWithTime(key, format string) log.Middleware {
	return func(ctx context.Context, e *entry.Entry, handler log.Logger) (int, error) {
		return handler(ctx, e.Add(field.FormatTime(key, format, time.Unix(math.MaxInt32, 0).In(time.UTC))))
	}
}

func ExampleFormatWithBracket() {
	ctx := context.Background()
	logger := log.New(log.WithFormat(log.FormatWithBracket()), log.WithStdout()).With(
		log.WithSource(10, filepath.Base),
		// log.WithTime(log.KeyTime, time.RFC3339),
		exampleWithTime(log.KeyTime, time.RFC3339),
		log.WithLevel(log.KeyLevel, level.Info),
	)

	logger.InfoKV(ctx, "imfo message", field.Int64("num", 42))

	serviceLogger := logger.With(log.WithName("service_name"))
	serviceLogger.Err(ctx, "error message")
	// Output:
	// 2038-01-19T03:14:07Z [info] writer_example_test.go:30 "imfo message" num=42
	// 2038-01-19T03:14:07Z [error][service_name] writer_example_test.go:33 "error message"
}
