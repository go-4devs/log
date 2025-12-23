package entry

import (
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

const (
	defaultCap = 5
)

type Option func(*Entry)

func WithCapacity(c int) Option {
	return func(e *Entry) {
		e.fields = make(field.Fields, 0, c+1)
	}
}

func WithFields(fields ...field.Field) Option {
	return func(e *Entry) {
		e.fields = fields
	}
}

func WithMessage(msg string) Option {
	return func(e *Entry) {
		e.format = msg
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
		fields: make(field.Fields, 0, defaultCap+1),
		level:  level.Debug,
		format: "",
		args:   make([]any, 0, defaultCap+1),
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
	fields field.Fields
}

func (e *Entry) Reset() {
	e.fields = e.fields[:0]
	e.args = e.args[:0]
	e.format = ""
}

func (e *Entry) Fields() field.Fields {
	return e.fields
}

// String implement stringer.
func (e *Entry) String() string {
	if e == nil {
		return ""
	}

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
	if e == nil {
		return level.Debug
	}

	return e.level
}

func (e *Entry) SetLevel(level level.Level) *Entry {
	if e == nil {
		return New().SetLevel(level)
	}

	e.level = level

	return e
}

func (e *Entry) SetMessage(msg string) *Entry {
	if e == nil {
		return New().SetMessage(msg)
	}

	e.format = msg

	return e
}

func (e *Entry) SetMessagef(format string, args ...any) *Entry {
	if e == nil {
		return New().SetMessagef(format, args...)
	}

	e.format = format
	e.args = append(e.args[:0], args...)

	return e
}

func (e *Entry) Add(fields ...field.Field) *Entry {
	if e == nil {
		return New(WithFields(fields...))
	}

	e.fields = e.fields.Append(fields...)

	return e
}

func (e *Entry) AddAny(key string, value any) *Entry {
	return e.Add(field.Any(key, value))
}

func (e *Entry) AddString(key, value string) *Entry {
	return e.Add(field.String(key, value))
}

func (e *Entry) Replace(key string, value field.Value) *Entry {
	has := false

	e.fields.Fields(func(f field.Field) bool {
		if f.Key == key {
			f.Value = value
			has = true

			return false
		}

		return true
	})

	if !has {
		e.AddAny(key, value)
	}

	return e
}
