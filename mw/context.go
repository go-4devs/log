package mw

import (
	"context"
	"fmt"

	"github.com/go-4devs/log"
)

// Context add field by context key
func Context(keys ...fmt.Stringer) log.Middleware {
	return func(ctx context.Context, level log.Level, msg string, fields log.Fields, handler log.Handler) {
		ctxFields := make(log.Fields, len(keys))
		for i, key := range keys {
			ctxFields[i] = log.Field{Key: key.String(), Value: ctx.Value(key)}
		}

		handler(ctx, level, msg, append(fields, ctxFields...))
	}
}
