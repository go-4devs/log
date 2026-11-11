package otel

import (
	"context"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/level"
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
		return levelFatal
	case level.Alert:
		return levelError3
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
	attrs := make([]attribute.KeyValue, 0, len(data.Fields())+levelFields)

	lvl := levels(data.Level())
	attrs = append(attrs,
		attribute.String(fieldSeverityText, lvl.String()),
		attribute.Int(fieldSeverityNumber, int(lvl)),
	)

	for _, attr := range data.Fields() {
		attrs = append(attrs, attribute.String(attr.Key, attr.Value.String()))
	}

	span.AddEvent(data.Message(), trace.WithAttributes(attrs...))
}
