package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"reflect"
	"slices"
	"sync"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/internal/buffer"
)

// Keys for built-in attributes.
const (
	// KeyTime is the key for the time attribute.
	KeyTime = "time"
	// KeyLevel is the key for the level name.
	KeyLevel = "level"
	// KeyMessage is the key for the log message.
	KeyMessage = "msg"
	// KeySource is the key for the caller location.
	KeySource = "source"
	// KeyName is the key for the logger name.
	KeyName = "name"
)

// WithWriter sets the output writer.
func WithWriter(w io.Writer) func(*option) {
	return func(o *option) {
		o.out = w
	}
}

// WithStdout writes logs to standard output.
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

// New creates standard logger.
func New(opts ...func(*option)) Logger {
	log := option{
		format: FormatString(field.NewEncoderText()),
		out:    os.Stderr,
	}

	for _, opt := range opts {
		opt(&log)
	}

	mu := mutexFor(log.out)

	return func(_ context.Context, entry *entry.Entry) (int, error) {
		mu.Lock()
		defer mu.Unlock()

		return log.format(log.out, entry)
	}
}

// shortcut: pointer-writer locks live for the process, drop the map if writers are created per line
//
//nolint:gochecknoglobals
var writerLocks sync.Map

func mutexFor(w io.Writer) *sync.Mutex {
	value := reflect.ValueOf(w)
	switch value.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Chan, reflect.Func, reflect.Slice, reflect.UnsafePointer:
		key := value.Pointer()
		if key == 0 {
			break
		}

		if found, ok := writerLocks.Load(key); ok {
			mu, _ := found.(*sync.Mutex)

			return mu
		}

		created := &sync.Mutex{}
		actual, _ := writerLocks.LoadOrStore(key, created)
		mu, _ := actual.(*sync.Mutex)

		return mu
	default:
	}

	return &sync.Mutex{}
}

// Encoder appends one attribute or value.
type Encoder interface {
	AppendValue(dst []byte, val slog.Value) []byte
	AppendField(dst []byte, val slog.Attr) []byte
}

type bracketHead struct {
	timeField   slog.Attr
	nameField   slog.Attr
	sourceField slog.Attr
	hasTime     bool
	hasName     bool
	hasSource   bool
	rest        []slog.Attr
}

// FormatWithBracket writes the level and message in a bracket layout.
func FormatWithBracket(enc Encoder) func(io.Writer, *entry.Entry) (int, error) {
	return func(w io.Writer, data *entry.Entry) (int, error) {
		buf := buffer.New()

		defer func() {
			buf.Free()
		}()

		head := bracketHeadFrom(data.Fields())

		if head.hasTime {
			*buf = enc.AppendValue(*buf, head.timeField.Value)
			_, _ = buf.WriteString(" ")
		}

		_, _ = buf.WriteString("[")
		*buf = enc.AppendValue(*buf, slog.StringValue(data.Level().String()))
		_, _ = buf.WriteString("]")

		if head.hasName {
			_, _ = buf.WriteString("[")
			*buf = enc.AppendValue(*buf, head.nameField.Value)
			_, _ = buf.WriteString("]")
		}

		if head.hasSource {
			_, _ = buf.WriteString(" ")
			*buf = enc.AppendValue(*buf, head.sourceField.Value)
			_, _ = buf.WriteString(" ")
		}

		*buf = enc.AppendValue(*buf, slog.StringValue(data.Message()))

		for _, item := range head.rest {
			*buf = enc.AppendField(*buf, item)
		}

		_, _ = buf.WriteString("\n")

		n, err := w.Write(*buf)
		if err != nil {
			return 0, fmt.Errorf("format text:%w", err)
		}

		return n, nil
	}
}

// FormatString writes the message and attributes as text.
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

// FormatJSON writes one JSON object.
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

func isKey(key string, keys ...string) bool {
	return slices.Contains(keys, key)
}

func bracketHeadFrom(fields []slog.Attr) bracketHead {
	var head bracketHead

	for _, item := range fields {
		switch {
		case !head.hasTime && item.Key == KeyTime:
			head.timeField, head.hasTime = item, true
		case !head.hasName && item.Key == KeyName:
			head.nameField, head.hasName = item, true
		case !head.hasSource && item.Key == KeySource:
			head.sourceField, head.hasSource = item, true
		case !isKey(item.Key, KeyTime, KeySource, KeyName, KeyLevel):
			head.rest = append(head.rest, item)
		}
	}

	return head
}
