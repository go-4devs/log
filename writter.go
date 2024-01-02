package log

import (
	"context"
	"fmt"
	"io"
	"os"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/internal/buffer"
)

// Keys for "built-in" attributes.
const (
	// TimeKey is the key used by the built-in handlers for the time
	// when the log method is called. The associated Value is a [time.Time].
	KeyTime = "time"
	// LevelKey is the key used by the built-in handlers for the level
	// of the log call. The associated value is a [Level].
	KeyLevel = "level"
	// MessageKey is the key used by the built-in handlers for the
	// message of the log call. The associated value is a string.
	KeyMessage = "msg"
	// SourceKey is the key used by the built-in handlers for the source file
	// and line of the log call. The associated value is a string.
	KeySource = "source"
)

func WithWriter(w io.Writer) func(*option) {
	return func(o *option) {
		o.out = w
	}
}

func WithStdout() func(*option) {
	return func(o *option) {
		o.out = os.Stdout
	}
}

// WithStringFormat sets format as simple string.
func WithStringFormat() func(*option) {
	return func(o *option) {
		o.format = formatText()
	}
}

// WithJSONFormat sets json output format.
func WithJSONFormat() func(*option) {
	return func(o *option) {
		o.format = formatJSON()
	}
}

type option struct {
	format func(io.Writer, *entry.Entry) (int, error)
	out    io.Writer
}

// New creates standart logger.
func New(opts ...func(*option)) Logger {
	log := option{
		format: formatText(),
		out:    os.Stderr,
	}

	for _, opt := range opts {
		opt(&log)
	}

	return func(_ context.Context, entry *entry.Entry) (int, error) {
		return log.format(log.out, entry)
	}
}

func formatText() func(io.Writer, *entry.Entry) (int, error) {
	enc := field.NewEncoderText()

	return func(w io.Writer, entry *entry.Entry) (int, error) {
		buf := buffer.New()
		defer func() {
			buf.Free()
		}()

		*buf = enc.AppendField(*buf, field.String(KeyMessage, entry.Message()))

		for _, field := range entry.Fields() {
			*buf = enc.AppendField(*buf, field)
		}

		_, _ = buf.WriteString("\n")

		n, err := w.Write(*buf)
		if err != nil {
			return 0, fmt.Errorf("format text:%w", err)
		}

		return n, nil
	}
}

func formatJSON() func(w io.Writer, entry *entry.Entry) (int, error) {
	enc := field.NewEncoderJSON()

	return func(w io.Writer, entry *entry.Entry) (int, error) {
		buf := buffer.New()
		defer func() {
			buf.Free()
		}()

		_, _ = buf.WriteString("{")
		*buf = enc.AppendField(*buf, field.String(KeyMessage, entry.Message()))

		for _, field := range entry.Fields() {
			*buf = enc.AppendField(*buf, field)
		}

		_, _ = buf.WriteString("}")
		_, _ = buf.WriteString("\n")

		n, err := w.Write(*buf)
		if err != nil {
			return 0, fmt.Errorf("format json:%w", err)
		}

		return n, nil
	}
}
