package logrus

import (
	"context"

	"github.com/go-4devs/log"
	lrg "github.com/sirupsen/logrus"
)

func New(logger *lrg.Logger) log.Handler {
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
	return func(ctx context.Context, level log.Level, msg string, fields log.Fields) {
		lf := make(lrg.Fields, len(fields))
		for _, f := range fields {
			lf[f.Key] = f.Value
		}
		logger.WithFields(lf).Log(levels[level], msg)
	}
}
