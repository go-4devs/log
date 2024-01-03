package logrus

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

// Standard create new standart logrus handler.
// Deprecated: delete after 0.7.0
func Standard() log.Logger {
	return New(logrus.StandardLogger())
}

// New create new logrus handler.
// Deprecated: delete after 0.7.0
func New(log *logrus.Logger) log.Logger {
	return func(ctx context.Context, data *entry.Entry) (int, error) {
		lrgFields := make(logrus.Fields, data.Fields().Len())
		data.Fields().Fields(func(f field.Field) bool {
			lrgFields[f.Key] = f.Value.Any()

			return true
		})

		entry := log.WithContext(ctx).WithFields(lrgFields)

		switch data.Level() {
		case level.Emergency:
			entry.Panic(data.Message())
		case level.Alert:
			entry.Fatal(data.Message())
		case level.Critical, level.Error:
			entry.Error(data.Message())
		case level.Warning:
			entry.Warn(data.Message())
		case level.Notice, level.Info:
			entry.Info(data.Message())
		case level.Debug:
			entry.Debug(data.Message())
		}

		return 0, nil
	}
}
