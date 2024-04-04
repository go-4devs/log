package main

import (
	"context"
	"time"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/field"
)

func main() {
	ctx := context.Background()

	log.DebugKV(ctx, "debug message")
	log.ErrKV(ctx, "error message")
	log.Errf(ctx, "format error message:%v", 42)
	log.Err(ctx, "error message", 42)
	service(ctx, log.Log())

	logger := log.New(log.WithJSONFormat()).With(log.WithSource(10, log.TrimPath), log.WithTime(log.KeyTime, time.RFC3339))
	logger.AlertKV(ctx, "alert message new logger", field.String("string", "value"))
	service(ctx, logger)

	strLogger := log.New(log.WithFormat(log.FormatWithBracket(field.NewEncoderText()))).With(log.WithSource(10, log.TrimPath), log.WithTime(log.KeyTime, time.RFC3339))
	strLogger.AlertKV(ctx, "alert message new txt logger", field.String("string", "value"))
	service(ctx, strLogger)
}

func service(ctx context.Context, logger log.Logger) {
	logger = logger.With(log.WithName("service"))
	logger.WarnKV(ctx, "warn service message")
	otherService(ctx, logger)
}

func otherService(ctx context.Context, logger log.Logger) {
	logger = logger.With(log.WithName("other_service"))
	logger.WarnKV(ctx, "warn other service message")
}
