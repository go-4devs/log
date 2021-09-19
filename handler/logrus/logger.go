package logrus

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/level"
)

// Option configure logger.
type Option func(*logger)

// WithLevel sets callback level to log level.
func WithLevel(level level.Level, c func(*logrus.Entry, string)) Option {
	return func(l *logger) {
		l.levels[level] = c
	}
}

// WithLogrus sets logrus logger.
func WithLogrus(logrus *logrus.Logger) Option {
	return func(l *logger) {
		l.logrus = logrus
	}
}

// New create new logrus handler.
func New(opts ...Option) log.Logger {
	log := logger{
		logrus: logrus.StandardLogger(),
		levels: map[level.Level]func(*logrus.Entry, string){
			level.Emergency: panicLog,
			level.Alert:     fatalLog,
			level.Critical:  errorLog,
			level.Error:     errorLog,
			level.Warning:   warnLog,
			level.Notice:    infoLog,
			level.Info:      infoLog,
			level.Debug:     debugLog,
		},
	}

	for _, o := range opts {
		o(&log)
	}

	return log.log
}

type logger struct {
	levels map[level.Level]func(l *logrus.Entry, msg string)
	logrus *logrus.Logger
}

func (l *logger) log(ctx context.Context, e *entry.Entry) (int, error) {
	lrgFields := make(logrus.Fields, e.Fields().Len())
	for _, field := range e.Fields() {
		lrgFields[string(field.Key())] = field.AsInterface()
	}

	l.levels[e.Level()](l.logrus.WithFields(lrgFields), e.Message())

	return 0, nil
}

func panicLog(e *logrus.Entry, msg string) {
	e.Panic(msg)
}

func fatalLog(e *logrus.Entry, msg string) {
	e.Fatal(msg)
}

func errorLog(e *logrus.Entry, msg string) {
	e.Error(msg)
}

func warnLog(e *logrus.Entry, msg string) {
	e.Warn(msg)
}

func infoLog(e *logrus.Entry, msg string) {
	e.Info(msg)
}

func debugLog(e *logrus.Entry, msg string) {
	e.Debug(msg)
}
