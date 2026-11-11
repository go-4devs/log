package entry

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
)

const (
	defaultCap = 5
)

type Option func(*Entry)

func WithCapacity(c int) Option {
	return func(e *Entry) {
		e.fields = make([]slog.Attr, 0, c+1)
	}
}

func WithFields(fields ...slog.Attr) Option {
	return func(e *Entry) {
		e.fields = fields
	}
}

func WithMessage(msg string) Option {
	return func(e *Entry) {
		e.SetMessage(msg)
	}
}

func WithMessagef(format string, args ...any) Option {
	return func(e *Entry) {
		e.format = format
		e.args = args
	}
}

func WithLevel(lvl level.Level) Option {
	return func(e *Entry) {
		e.level = lvl
	}
}

func New(opts ...Option) *Entry {
	entry := &Entry{
		fields: make([]slog.Attr, 0, defaultCap+1),
		level:  level.Debug,
		format: "",
		args:   nil,
	}

	for _, opt := range opts {
		opt(entry)
	}

	return entry
}

// Entry slice field.
type Entry struct {
	format string
	args   []any
	level  level.Level
	fields []slog.Attr
}

func (e *Entry) Reset() {
	e.fields = e.fields[:0]
	e.args = e.args[:0]
	e.format = ""
	e.level = level.Debug
}

func (e *Entry) Fields() []slog.Attr {
	return e.fields
}

// Clone returns a copy with its own fields and args.
func (e *Entry) Clone() *Entry {
	clone := *e
	clone.fields = slices.Clone(e.fields)
	clone.args = slices.Clone(e.args)

	return &clone
}

// String implement stringer.
func (e *Entry) String() string {
	str := make([]string, len(e.fields)+1)
	str[0] = e.Message()

	for i, field := range e.fields {
		str[i+1] = field.String()
	}

	return strings.Join(str, " ")
}

func (e *Entry) Message() string {
	switch {
	case len(e.args) > 0 && e.format != "":
		return fmt.Sprintf(e.format, e.args...)
	case len(e.args) > 0:
		return fmt.Sprint(e.args...)
	default:
		return e.format
	}
}

func (e *Entry) Level() level.Level {
	return e.level
}

func (e *Entry) SetLevel(level level.Level) *Entry {
	e.level = level

	return e
}

func (e *Entry) SetMessage(msg string) *Entry {
	e.format = msg
	e.args = nil

	return e
}

func (e *Entry) SetMessagef(format string, args ...any) *Entry {
	e.format = format
	e.args = append(e.args[:0], args...)

	return e
}

func (e *Entry) Add(fields ...slog.Attr) *Entry {
	e.fields = append(e.fields, fields...)

	return e
}

func (e *Entry) AddAny(key string, value any) *Entry {
	return e.Add(field.Any(key, value))
}

func (e *Entry) AddString(key, value string) *Entry {
	return e.Add(field.String(key, value))
}

func (e *Entry) Replace(key string, value slog.Value) *Entry {
	attr := field.Any(key, value)

	for i := range e.fields {
		if e.fields[i].Key == key {
			e.fields[i] = attr

			return e
		}
	}

	return e.Add(attr)
}
