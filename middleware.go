package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

// Middleware handle.
type Middleware func(ctx context.Context, level Level, msg string, fields Fields, handler Logger)

// With add middleware to logger.
func With(logger Logger, mw ...Middleware) Logger {
	switch len(mw) {
	case 0:
		return logger
	case 1:
		return func(ctx context.Context, level Level, msg string, fields Fields) {
			mw[0](ctx, level, msg, fields, logger)
		}
	}

	lastI := len(mw) - 1

	return func(ctx context.Context, level Level, msg string, fields Fields) {
		var (
			chainHandler func(ctx context.Context, level Level, msg string, fields Fields)
			curI         int
		)

		chainHandler = func(currentCtx context.Context, currentLevel Level, currentMsg string, currentFields Fields) {
			if curI == lastI {
				logger(currentCtx, currentLevel, currentMsg, currentFields)
				return
			}
			curI++
			mw[curI](currentCtx, currentLevel, currentMsg, currentFields, chainHandler)
			curI--
		}

		mw[0](ctx, level, msg, fields, chainHandler)
	}
}

// WithLevel sets log level.
func WithLevel(lvl Level) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		if level <= lvl {
			handler(ctx, level, msg, append(fields, Field{Key: "level", Value: level}))
		}
	}
}

// KeyValue add field by const key value.
func KeyValue(key string, value interface{}) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		handler(ctx, level, msg, append(fields, Field{Key: key, Value: value}))
	}
}

// GoVersion add field by go version.
func GoVersion(key string) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		handler(ctx, level, msg, append(fields, Field{Key: key, Value: runtime.Version()}))
	}
}

// WithContext add field by context key.
func WithContextValue(keys ...fmt.Stringer) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		ctxFields := make(Fields, len(keys))
		for i, key := range keys {
			ctxFields[i] = Field{Key: key.String(), Value: ctx.Value(key)}
		}

		handler(ctx, level, msg, append(fields, ctxFields...))
	}
}

// WithCaller adds called file.
func WithCaller(calldepth int, short bool) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		_, file, line, ok := runtime.Caller(calldepth)
		if !ok {
			file, line = "???", 0
		}

		if short && ok {
			file = filepath.Base(file)
		}

		handler(ctx, level, msg, append(fields, NewField("caller", fmt.Sprint(file, ":", line))))
	}
}

// WithTime adds time.
func WithTime(format string) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		handler(ctx, level, msg, append(fields, NewField("time", time.Now().Format(format))))
	}
}

// WithMetrics adds handle metrics.
func WithMetrics(metrics func(level Level)) Middleware {
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Logger) {
		go metrics(level)
		handler(ctx, level, msg, fields)
	}
}
