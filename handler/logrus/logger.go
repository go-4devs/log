package logrus

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/level"
)

// Standard create new standart logrus handler.
func Standard() log.Logger {
	return New(logrus.StandardLogger())
}

// New create new logrus handler.
func New(log *logrus.Logger) log.Logger {
	return func(ctx context.Context, e *entry.Entry) (int, error) {
		lrgFields := make(logrus.Fields, e.Fields().Len())
		for _, field := range e.Fields() {
			lrgFields[string(field.Key())] = field.AsInterface()
		}

		entry := log.WithContext(ctx).WithFields(lrgFields)

		switch e.Level() {
		case level.Emergency:
			entry.Panic(e.Message())
		case level.Alert:
			entry.Fatal(e.Message())
		case level.Critical, level.Error:
			entry.Error(e.Message())
		case level.Warning:
			entry.Warn(e.Message())
		case level.Notice, level.Info:
			entry.Info(e.Message())
		case level.Debug:
			entry.Debug(e.Message())
		}

		return 0, nil
	}
}
