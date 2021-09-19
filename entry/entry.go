package entry

import (
	"strings"

	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

const (
	defaultCap = 5
)

type Option func(*Entry)

func WithCapacity(cap int) Option {
	return func(e *Entry) {
		e.fields = make(field.Fields, 0, cap+1)
	}
}

func WithFields(fields ...field.Field) Option {
	return func(e *Entry) {
		e.fields = fields
	}
}

func WithMessage(msg string) Option {
	return func(e *Entry) {
		e.msg = msg
	}
}

func WithLevel(lvl level.Level) Option {
	return func(e *Entry) {
		e.level = lvl
	}
}

func New(opts ...Option) *Entry {
	e := &Entry{
		fields: make(field.Fields, 0, defaultCap+1),
		level:  level.Debug,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Entry slice field.
type Entry struct {
	msg    string
	level  level.Level
	fields field.Fields
}

func (e *Entry) Reset() {
	e.fields = e.fields[:0]
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
	str[0] = e.msg

	for i, field := range e.fields {
		str[i+1] = field.String()
	}

	return strings.Join(str, " ")
}

func (e *Entry) Message() string {
	return e.msg
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

	e.msg = msg

	return e
}

func (e *Entry) Add(fields ...field.Field) *Entry {
	if e == nil {
		return New(WithFields(fields...))
	}

	e.fields = e.fields.Append(fields...)

	return e
}

func (e *Entry) AddAny(key string, value interface{}) *Entry {
	return e.Add(field.Any(key, value))
}

func (e *Entry) AddString(key, value string) *Entry {
	return e.Add(field.String(key, value))
}
