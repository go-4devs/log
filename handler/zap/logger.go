package zap

import (
	"context"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/level"
	"go.uber.org/zap"
)

// Option configure logger.
type Option func(*logger)

// WithLevel sets level logged message.
func WithLevel(level level.Level, f func(z *zap.Logger, msg string, fields ...zap.Field)) Option {
	return func(l *logger) {
		l.levels[level] = f
	}
}

// WithZap sets zap logger.
func WithZap(z *zap.Logger) Option {
	return func(l *logger) {
		l.zap = z
	}
}

// New create handler by zap logger.
func New(opts ...Option) log.Logger {
	z, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log := logger{
		zap: z,
		levels: map[level.Level]func(z *zap.Logger, msg string, fields ...zap.Field){
			level.Emergency: fatalLog,
			level.Alert:     panicLog,
			level.Critical:  errorLog,
			level.Error:     errorLog,
			level.Warning:   warnLog,
			level.Notice:    infoLog,
			level.Info:      infoLog,
			level.Debug:     debugLog,
		},
	}

	for _, opt := range opts {
		opt(&log)
	}

	return log.log
}

type logger struct {
	zap    *zap.Logger
	levels map[level.Level]func(z *zap.Logger, msg string, fields ...zap.Field)
}

func (l *logger) log(ctx context.Context, e *entry.Entry) (int, error) {
	zf := make([]zap.Field, e.Fields().Len())
	for i, field := range e.Fields() {
		zf[i] = zap.Any(string(field.Key()), field.AsInterface())
	}

	l.levels[e.Level()](l.zap, e.Message(), zf...)

	return 0, nil
}

func panicLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Panic(msg, fields...)
}

func fatalLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Fatal(msg, fields...)
}

func errorLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Error(msg, fields...)
}

func warnLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Warn(msg, fields...)
}

func infoLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Info(msg, fields...)
}

func debugLog(z *zap.Logger, msg string, fields ...zap.Field) {
	z.Debug(msg, fields...)
}
