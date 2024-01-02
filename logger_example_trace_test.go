package log_test

import (
	"context"
	"fmt"
	"io"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/handler/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func ExampleNew_withTrace() {
	logger := log.New(log.WithStdout()).With(otel.Middleware())

	sctx, span := startSpan(ctx)

	logger.Err(sctx, "log logrus")
	logger.ErrKV(sctx, "log logrus kv", field.Int("int", 42))
	logger.ErrKVs(sctx, "log logrus kv sugar", "err", io.EOF)

	span.End()

	// Output:
	// msg="log logrus"
	// msg="log logrus kv" int=42
	// msg="log logrus kv sugar" err=EOF
	// event: log logrus, SeverityText = ERROR, SeverityNumber = 17
	// event: log logrus kv, SeverityText = ERROR, SeverityNumber = 17, int = 42
	// event: log logrus kv sugar, SeverityText = ERROR, SeverityNumber = 17, err = EOF
}

func startSpan(ctx context.Context) (context.Context, trace.Span) {
	tp := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter{}))

	return tp.Tracer("logger").Start(ctx, "operation")
}

type exporter struct{}

func (e exporter) Shutdown(_ context.Context) error {
	return nil
}

func (e exporter) ExportSpans(_ context.Context, spanData []sdktrace.ReadOnlySpan) error {
	for _, data := range spanData {
		for _, events := range data.Events() {
			fmt.Print("event: ", events.Name)

			for _, attr := range events.Attributes {
				fmt.Printf(", %v = %v", attr.Key, attr.Value.AsInterface())
			}

			fmt.Print("\n")
		}
	}

	return nil
}
