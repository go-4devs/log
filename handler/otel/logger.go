package otel

import (
	"context"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
)

// Deprecated: delete after 0.7.0
func New() log.Logger {
	return func(ctx context.Context, e *entry.Entry) (int, error) {
		addEvent(ctx, e)

		return 0, nil
	}
}
