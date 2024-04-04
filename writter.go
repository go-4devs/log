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
	// KeyName logger name.
	KeyName = "name"
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
	return WithFormat(FormatString(field.NewEncoderText()))
}

// WithJSONFormat sets json output format.
func WithJSONFormat() func(*option) {
	return WithFormat(FormatJSON(field.NewEncoderJSON()))
}

// WithFormat sets custom output format.
func WithFormat(format func(io.Writer, *entry.Entry) (int, error)) func(*option) {
	return func(o *option) {
		o.format = format
	}
}

type option struct {
	format func(io.Writer, *entry.Entry) (int, error)
	out    io.Writer
}

// New creates standart logger.
func New(opts ...func(*option)) Logger {
	log := option{
		format: FormatString(field.NewEncoderText()),
		out:    os.Stderr,
	}

	for _, opt := range opts {
		opt(&log)
	}

	return func(_ context.Context, entry *entry.Entry) (int, error) {
		return log.format(log.out, entry)
	}
}

type Encoder interface {
	AppendValue(dst []byte, val field.Value) []byte
	AppendField(dst []byte, val field.Field) []byte
}

func FormatWithBracket(enc Encoder) func(io.Writer, *entry.Entry) (int, error) {
	appendValue := func(buf *buffer.Buffer, data field.Fields, key, prefix, suffix string) *buffer.Buffer {
		data.Fields(
			func(f field.Field) bool {
				if f.IsKey(key) {
					_, _ = buf.WriteString(prefix)
					*buf = enc.AppendValue(*buf, f.Value)
					_, _ = buf.WriteString(suffix)

					return false
				}

				return true
			})

		return buf
	}

	return func(w io.Writer, data *entry.Entry) (int, error) {
		buf := buffer.New()
		defer func() {
			buf.Free()
		}()

		fields := data.Fields()
		buf = appendValue(buf, fields, KeyTime, "", " ")
		_, _ = buf.WriteString("[")
		*buf = enc.AppendValue(*buf, field.StringValue(data.Level().String()))
		_, _ = buf.WriteString("]")
		buf = appendValue(buf, fields, KeyName, "[", "]")
		buf = appendValue(buf, fields, KeySource, " ", " ")
		*buf = enc.AppendValue(*buf, field.StringValue(data.Message()))

		fields.Fields(func(f field.Field) bool {
			if !f.IsKey(KeyTime, KeySource, KeyName, KeyLevel) {
				*buf = enc.AppendField(*buf, f)
			}

			return true
		})

		_, _ = buf.WriteString("\n")

		n, err := w.Write(*buf)
		if err != nil {
			return 0, fmt.Errorf("format text:%w", err)
		}

		return n, nil
	}
}

func FormatString(enc Encoder) func(io.Writer, *entry.Entry) (int, error) {
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

func FormatJSON(enc Encoder) func(w io.Writer, entry *entry.Entry) (int, error) {
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
