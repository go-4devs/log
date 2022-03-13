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
func New(logger *zap.Logger) log.Logger {
	return func(ctx context.Context, data *entry.Entry) (int, error) {
		zf := make([]zap.Field, data.Fields().Len())
		for i, field := range data.Fields() {
			zf[i] = zap.Any(string(field.Key()), field.AsInterface())
		}

		switch data.Level() {
		case level.Emergency:
			logger.Fatal(data.Message(), zf...)
		case level.Alert:
			logger.Panic(data.Message(), zf...)
		case level.Critical, level.Error:
			logger.Error(data.Message(), zf...)
		case level.Warning:
			logger.Warn(data.Message(), zf...)
		case level.Notice, level.Info:
			logger.Info(data.Message(), zf...)
		case level.Debug:
			logger.Debug(data.Message(), zf...)
		}

		return 0, nil
	}
}
