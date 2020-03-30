package mw

import (
	"context"
	"runtime"

	"github.com/go-4devs/log"
)

// GoVersion add field by go version
func GoVersion(key string) log.Middleware {
	return func(ctx context.Context, level log.Level, msg string, fields log.Fields, handler log.Handler) {
		handler(ctx, level, msg, append(fields, log.Field{Key: key, Value: runtime.Version()}))
	}
}
