package zap

import (
	"context"
	"time"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Deprecated: delete after 0.7.0
func Nop() log.Logger {
	return New(zap.NewNop())
}

// Deprecated: delete after 0.7.0
func Example(options ...zap.Option) log.Logger {
	return New(zap.NewExample(options...))
}

// Deprecated: delete after 0.7.0
func Production(options ...zap.Option) log.Logger {
	z, err := zap.NewProduction(options...)
	if err != nil {
		panic(err)
	}

	return New(z)
}

// Deprecated: delete after 0.7.0
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
		zf := make([]zap.Field, 0, len(data.Fields()))
		for _, attr := range data.Fields() {
			zf = append(zf, zap.Any(attr.Key, field.ValueAny(attr.Value)))
		}

		switch data.Level() {
		case level.Emergency:
			logAt(logger, zapcore.PanicLevel, data.Message(), zf)
		case level.Alert:
			logAt(logger, zapcore.FatalLevel, data.Message(), zf)
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

// logAt writes at lvl without zap's panic or exit hooks.
func logAt(logger *zap.Logger, lvl zapcore.Level, msg string, fields []zap.Field) {
	ce := logger.Core().Check(zapcore.Entry{
		LoggerName: logger.Name(),
		Time:       time.Now(),
		Level:      lvl,
		Message:    msg,
	}, nil)
	if ce != nil {
		ce.Write(fields...)
	}
}
