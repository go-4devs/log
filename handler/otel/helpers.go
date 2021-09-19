package otel

import (
	"context"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/level"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/label"
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

func addEvent(ctx context.Context, e *entry.Entry) {
	span := trace.SpanFromContext(ctx)
	attrs := make([]label.KeyValue, 0, e.Fields().Len()+levelFields)

	lvl := levels(e.Level())
	attrs = append(attrs,
		label.String(fieldSeverityText, lvl.String()),
		label.Int(fieldSeverityNumber, int(lvl)),
	)

	for _, field := range e.Fields() {
		attrs = append(attrs, label.String(string(field.Key()), field.Value().String()))
	}

	span.AddEvent(ctx, e.Message(), attrs...)
}
