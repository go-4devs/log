package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

// Middleware handle.
type Middleware func(ctx context.Context, e *entry.Entry, handler Logger) (int, error)

// With add middleware to logger.
func With(logger Logger, mw ...Middleware) Logger {
	switch len(mw) {
	case 0:
		return logger
	case 1:
		return func(ctx context.Context, entry *entry.Entry) (int, error) {
			return mw[0](ctx, entry, logger)
		}
	}

	lastI := len(mw) - 1

	return func(ctx context.Context, data *entry.Entry) (int, error) {
		var (
			chainHandler func(context.Context, *entry.Entry) (int, error)
			curI         int
		)

		chainHandler = func(currentCtx context.Context, currentEntry *entry.Entry) (int, error) {
			if curI == lastI {
				return logger(currentCtx, currentEntry)
			}

			curI++
			n, err := mw[curI](currentCtx, currentEntry, chainHandler)
			curI--

			return n, err
		}

		return mw[0](ctx, data, chainHandler)
	}
}

// WithLevel sets log level.
func WithLevel(key string, lvl level.Level) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		if e.Level().Enabled(lvl) {
			return handler(ctx, e.AddString(key, e.Level().String()))
		}

		return 0, nil
	}
}

// KeyValue add field by const key value.
func KeyValue(key string, value any) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		return handler(ctx, e.AddAny(key, value))
	}
}

// GoVersion add field by go version.
func GoVersion(key string) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		return handler(ctx, e.AddString(key, runtime.Version()))
	}
}

// WithContextValue add field by context key.
func WithContextValue(keys ...fmt.Stringer) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		for _, key := range keys {
			e = e.AddAny(key.String(), ctx.Value(key))
		}

		return handler(ctx, e)
	}
}

func WithName(name string) Middleware {
	return func(ctx context.Context, data *entry.Entry, handler Logger) (int, error) {
		return handler(ctx, data.Replace(KeyName, field.StringValue(name)))
	}
}

// WithCaller adds called file.
//
// Deprecated: use WithSource.
func WithCaller(key string, depth int, full bool) Middleware {
	const offset = 2

	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		return handler(ctx, e.AddString(key, entry.Caller(depth*offset, full)))
	}
}

// WithTime adds time.
func WithTime(key, format string) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		return handler(ctx, e.Add(field.FormatTime(key, format, time.Now())))
	}
}

// WithMetrics adds handle metrics.
func WithMetrics(metrics func(level level.Level)) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		go metrics(e.Level())

		return handler(ctx, e)
	}
}

// WithExit exit by level.
func WithExit(level level.Level) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		n, err := handler(ctx, e)

		if e.Level().Is(level) {
			os.Exit(1)
		}

		return n, err
	}
}

// WithPanic panic by level.
func WithPanic(level level.Level) Middleware {
	return func(ctx context.Context, e *entry.Entry, handler Logger) (int, error) {
		n, err := handler(ctx, e)

		if e.Level().Is(level) {
			panic(e.String())
		}

		return n, err
	}
}
