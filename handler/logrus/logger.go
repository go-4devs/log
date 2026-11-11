package logrus

import (
	"context"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
	"github.com/sirupsen/logrus"
)

// Standard create new standard logrus handler.
// Deprecated: delete after 0.7.0
func Standard() log.Logger {
	return New(logrus.StandardLogger())
}

// New create new logrus handler.
// Deprecated: delete after 0.7.0
func New(logger *logrus.Logger) log.Logger {
	return func(ctx context.Context, data *entry.Entry) (int, error) {
		lrgFields := make(logrus.Fields, data.Fields().Len())
		data.Fields().Fields(func(f field.Field) bool {
			lrgFields[f.Key] = f.Value.Any()

			return true
		})

		row := logger.WithContext(ctx).WithFields(lrgFields)

		switch data.Level() {
		case level.Emergency, level.Alert:
			row.WithField(log.KeyLevel, data.Level().String()).Log(logrus.ErrorLevel, data.Message())
		case level.Critical, level.Error:
			row.Error(data.Message())
		case level.Warning:
			row.Warn(data.Message())
		case level.Notice, level.Info:
			row.Info(data.Message())
		case level.Debug:
			row.Debug(data.Message())
		}

		return 0, nil
	}
}
