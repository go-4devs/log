package mw

import (
	"context"

	"github.com/go-4devs/log"
)

// KeyValue add field by const key value
func KeyValue(key string, value interface{}) log.Middleware {
	return func(ctx context.Context, level log.Level, msg string, fields log.Fields, handler log.Handler) {
		handler(ctx, level, msg, append(fields, log.Field{Key: key, Value: value}))
	}
}
