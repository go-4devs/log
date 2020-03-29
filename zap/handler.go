package zap

import (
	"context"

	"github.com/go-4devs/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Levels map levels
type Levels map[log.Level]zapcore.Level

// Option configure levels
type Option func(Levels)

// WithLevel sete level logged message
func WithLevel(level log.Level, zapLevel zapcore.Level) Option {
	return func(l Levels) {
		l[level] = zapLevel
	}
}

// New ceate handler by zap logger
func New(z *zap.Logger, opts ...Option) log.Handler {
	levels := map[log.Level]zapcore.Level{
		log.LevelEmergency: zapcore.FatalLevel,
		log.LevelAlert:     zapcore.PanicLevel,
		log.LevelCritical:  zapcore.ErrorLevel,
		log.LevelError:     zapcore.ErrorLevel,
		log.LevelWarning:   zapcore.WarnLevel,
		log.LevelNotice:    zapcore.InfoLevel,
		log.LevelInfo:      zapcore.InfoLevel,
		log.LevelDebug:     zapcore.DebugLevel,
	}

	for _, o := range opts {
		o(levels)
	}

	return func(ctx context.Context, level log.Level, msg string, fields log.Fields) {
		zf := make([]zap.Field, len(fields))
		for i, field := range fields {
			zf[i] = zap.Any(field.Key, field.Value)
		}

		if ce := z.Check(levels[level], msg); ce != nil {
			ce.Write(zf...)
		}
	}
}
