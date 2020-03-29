package log

import (
	"context"
	"log"
)

// NewStdHandler create handler for the log package
func NewStdHandler(logger *log.Logger, lvl Level) Handler {
	return func(ctx context.Context, level Level, msg string, fields Fields) {
		if lvl < level {
			return
		}

		switch level {
		case LevelEmergency, LevelAlert:
			logger.Fatal("level: ", level.String(), " msg:; ", msg, ";", fields)
		}

		logger.Print("level: ", level.String(), "; msg: ", msg, ";", fields)
	}
}
