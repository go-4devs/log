package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	calldepth = 3
)

// New creates standart logger.
func New(opts ...Option) Logger {
	logger := logger{
		format:    stringFormat,
		output:    log.Output,
		calldepth: calldepth,
	}

	for _, opt := range opts {
		opt(&logger)
	}

	return func(ctx context.Context, level Level, msg string, fields Fields) {
		_ = logger.output(logger.calldepth, logger.format(msg, fields))

		switch level {
		case LevelEmergency:
			panic(msg)
		case LevelAlert:
			os.Exit(1)
		default:
		}
	}
}

// Option configure log.
type Option func(*logger)

// Format sets formats output message.
type Format func(msg string, fields Fields) string

type logger struct {
	output    func(calldepth int, s string) error
	format    Format
	calldepth int
}

// WithWriter sets writer logger.
func WithWriter(writer io.Writer) Option {
	return func(l *logger) {
		l.output = log.New(writer, "", 0).Output
	}
}

// WithStdout sets logged to os.Stdout.
func WithStdout() Option {
	return func(l *logger) {
		l.output = log.New(os.Stdout, "", 0).Output
	}
}

// WithFormat sets format log.
func WithFormat(format Format) Option {
	return func(l *logger) {
		l.format = format
	}
}

// WithStringFormat sets format as simple string.
func WithStringFormat() Option {
	return func(l *logger) {
		l.format = stringFormat
	}
}

// WithJSONFormat sets json output format.
func WithJSONFormat() Option {
	return func(l *logger) {
		l.format = jsonFormat
	}
}

// WithCalldepth sets depth filename.
func WithCalldepth(calldepth int) Option {
	return func(l *logger) {
		l.calldepth = calldepth
	}
}

// WithLogger sets logger anf format.
func WithLogger(std *log.Logger, format Format) Option {
	return func(l *logger) {
		l.output = std.Output
		l.format = format
	}
}

func stringFormat(msg string, fields Fields) string {
	return fmt.Sprint("msg=\"", msg, "\" ", fields)
}

func jsonFormat(msg string, fields Fields) string {
	data := make(map[string]interface{}, len(fields)+1)
	data["msg"] = msg

	for _, field := range fields {
		data[field.Key] = field.Value
	}

	res, err := json.Marshal(data)
	if err != nil {
		return stringFormat(msg, append(fields, FieldError(err)))
	}

	return string(res)
}
