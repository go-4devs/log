package otel_test

import (
	"context"
	"io"
	"testing"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/handler/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func TestSeverityOrder(t *testing.T) {
	t.Parallel()

	exp := &recordingExporter{}
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exp))
	ctx, span := provider.Tracer("test").Start(context.Background(), "op")

	logger := log.New(log.WithWriter(io.Discard)).With(otel.Middleware())
	logger.Emerg(ctx, "emerg")
	logger.Alert(ctx, "alert")
	span.End()

	if len(exp.spans) != 1 {
		t.Fatalf("spans: %d", len(exp.spans))
	}

	text, number := severity(exp.spans[0], "emerg")
	if text != "FATAL" || number != 21 {
		t.Fatalf("emergency: %s %d", text, number)
	}

	text, number = severity(exp.spans[0], "alert")
	if text != "ERROR3" || number != 19 {
		t.Fatalf("alert: %s %d", text, number)
	}
}

func severity(span sdktrace.ReadOnlySpan, name string) (string, int64) {
	for _, event := range span.Events() {
		if event.Name != name {
			continue
		}

		var (
			text   string
			number int64
		)

		for _, attr := range event.Attributes {
			switch string(attr.Key) {
			case "SeverityText":
				text = attr.Value.AsString()
			case "SeverityNumber":
				number = attr.Value.AsInt64()
			}
		}

		return text, number
	}

	return "", 0
}

type recordingExporter struct {
	spans []sdktrace.ReadOnlySpan
}

func (e *recordingExporter) Shutdown(context.Context) error {
	return nil
}

func (e *recordingExporter) ExportSpans(_ context.Context, spans []sdktrace.ReadOnlySpan) error {
	e.spans = append(e.spans, spans...)

	return nil
}
