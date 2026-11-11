// Package otel records log entries as OpenTelemetry events.
package otel

import (
	"context"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
)

// New returns a logger that records an OpenTelemetry event.
//
// Deprecated: delete after 0.7.0
func New() log.Logger {
	return func(ctx context.Context, e *entry.Entry) (int, error) {
		addEvent(ctx, e)

		return 0, nil
	}
}
