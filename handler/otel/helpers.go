package otel

import (
	"context"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	fieldSeverityNumber = "SeverityNumber"
	fieldSeverityText   = "SeverityText"
	levelFields         = 2
)

func levels(lvl level.Level) Level {
	switch lvl {
	case level.Emergency:
		return levelError3
	case level.Alert:
		return levelFatal
	case level.Critical:
		return levelError2
	case level.Error:
		return levelError
	case level.Warning:
		return levelWarn
	case level.Notice:
		return levelInfo2
	case level.Info:
		return levelInfo
	case level.Debug:
		return levelDebug
	}

	return 0
}

func addEvent(ctx context.Context, data *entry.Entry) {
	span := trace.SpanFromContext(ctx)
	attrs := make([]attribute.KeyValue, 0, data.Fields().Len()+levelFields)

	lvl := levels(data.Level())
	attrs = append(attrs,
		attribute.String(fieldSeverityText, lvl.String()),
		attribute.Int(fieldSeverityNumber, int(lvl)),
	)

	data.Fields().Fields(func(f field.Field) bool {
		attrs = append(attrs, attribute.String(f.Key, f.Value.String()))

		return true
	})

	span.AddEvent(data.Message(), trace.WithAttributes(attrs...))
}
