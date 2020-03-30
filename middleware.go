package log

import (
	"context"
)

type Middleware func(ctx context.Context, level Level, msg string, fields Fields, handler Handler)

func NewHandler(handler Handler, mw ...Middleware) Handler {
	if len(mw) == 0 {
		return handler
	}
	m := ChainMiddlerware(mw...)
	return func(ctx context.Context, level Level, msg string, fields Fields) {
		m(ctx, level, msg, fields, handler)
	}
}

func ChainMiddlerware(mw ...Middleware) Middleware {
	lastI := len(mw) - 1
	return func(ctx context.Context, level Level, msg string, fields Fields, handler Handler) {
		var (
			chainHandler func(ctx context.Context, level Level, msg string, fields Fields)
			curI         int
		)

		chainHandler = func(currentCtx context.Context, currentLevel Level, currentMsg string, currentFields Fields) {
			if curI == lastI {
				handler(currentCtx, currentLevel, currentMsg, currentFields)
				return
			}
			curI++
			mw[curI](currentCtx, currentLevel, currentMsg, currentFields, chainHandler)
			curI--
		}

		mw[0](ctx, level, msg, fields, chainHandler)
	}
}
