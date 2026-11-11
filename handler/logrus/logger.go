// Package logrus writes log entries with a logrus logger.
package logrus

import (
	"context"

	"github.com/go-4devs/log"
	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
	"github.com/sirupsen/logrus"
)

// Standard returns a logger around the standard logrus logger.
//
// Deprecated: delete after 0.7.0
func Standard() log.Logger {
	return New(logrus.StandardLogger())
}

// New returns a logger that writes to logger.
//
// Deprecated: delete after 0.7.0
func New(logger *logrus.Logger) log.Logger {
	return func(ctx context.Context, data *entry.Entry) (int, error) {
		lrgFields := make(logrus.Fields, len(data.Fields()))
		for _, attr := range data.Fields() {
			lrgFields[attr.Key] = field.ValueAny(attr.Value)
		}

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
