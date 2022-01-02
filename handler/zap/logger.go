package zap

import (
	"context"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/level"
	"go.uber.org/zap"
)

func Nop() log.Logger {
	return New(zap.NewNop())
}

func Example(options ...zap.Option) log.Logger {
	return New(zap.NewExample(options...))
}

func Production(options ...zap.Option) log.Logger {
	z, err := zap.NewProduction(options...)
	if err != nil {
		panic(err)
	}

	return New(z)
}

func Development(options ...zap.Option) log.Logger {
	z, err := zap.NewDevelopment(options...)
	if err != nil {
		panic(err)
	}

	return New(z)
}

// New create handler by zap logger.
func New(z *zap.Logger) log.Logger {
	return func(ctx context.Context, e *entry.Entry) (int, error) {
		zf := make([]zap.Field, e.Fields().Len())
		for i, field := range e.Fields() {
			zf[i] = zap.Any(string(field.Key()), field.AsInterface())
		}

		switch e.Level() {
		case level.Emergency:
			z.Fatal(e.Message(), zf...)
		case level.Alert:
			z.Panic(e.Message(), zf...)
		case level.Critical, level.Error:
			z.Error(e.Message(), zf...)
		case level.Warning:
			z.Warn(e.Message(), zf...)
		case level.Notice, level.Info:
			z.Info(e.Message(), zf...)
		case level.Debug:
			z.Debug(e.Message(), zf...)
		}

		return 0, nil
	}
}
