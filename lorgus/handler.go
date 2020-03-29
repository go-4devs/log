package logrus

import (
	"context"

	"github.com/go-4devs/log"
	lrg "github.com/sirupsen/logrus"
)

// Levels maps
type Levels map[log.Level]lrg.Level

// Option configure levels
type Option func(Levels)

// WithLevel set lorgus level to log level
func WithLevel(level log.Level, loggusLevel lrg.Level) Option {
	return func(l Levels) {
		l[level] = loggusLevel
	}
}

// New create new lorgus handler
func New(logger *lrg.Logger, level log.Level, opts ...Option) log.Handler {
	levels := map[log.Level]lrg.Level{
		log.LevelEmergency: lrg.PanicLevel,
		log.LevelAlert:     lrg.FatalLevel,
		log.LevelCritical:  lrg.ErrorLevel,
		log.LevelError:     lrg.ErrorLevel,
		log.LevelWarning:   lrg.WarnLevel,
		log.LevelNotice:    lrg.InfoLevel,
		log.LevelInfo:      lrg.InfoLevel,
		log.LevelDebug:     lrg.DebugLevel,
	}

	for _, o := range opts {
		o(levels)
	}

	logger.SetLevel(levels[level])

	return func(ctx context.Context, level log.Level, msg string, fields log.Fields) {
		lf := make(lrg.Fields, len(fields))
		for _, f := range fields {
			lf[f.Key] = f.Value
		}

		logger.WithFields(lf).Log(levels[level], msg)
	}
}
