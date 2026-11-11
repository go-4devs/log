package otel

import (
	"context"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
)

// Deprecated: delete after 0.7.0
func Middleware() log.Middleware {
	return func(ctx context.Context, e *entry.Entry, handler log.Logger) (int, error) {
		addEvent(ctx, e)

		return handler(ctx, e)
	}
}
