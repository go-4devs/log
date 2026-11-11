// Package entry holds one log record: a level, a message, and slog attributes.
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

// Option configures an Entry created by New.
type Option func(*Entry)

// WithCapacity sets the attribute slice capacity.
func WithCapacity(c int) Option {
	return func(e *Entry) {
		e.fields = make([]slog.Attr, 0, c+1)
	}
}

// WithFields sets the attributes. The slice is not copied.
func WithFields(fields ...slog.Attr) Option {
	return func(e *Entry) {
		e.fields = fields
	}
}

// WithMessage sets the message and drops format arguments.
func WithMessage(msg string) Option {
	return func(e *Entry) {
		e.SetMessage(msg)
	}
}

// WithMessagef sets a format and arguments rendered by Message.
func WithMessagef(format string, args ...any) Option {
	return func(e *Entry) {
		e.format = format
		e.args = args
	}
}

// WithLevel sets the level.
func WithLevel(lvl level.Level) Option {
	return func(e *Entry) {
		e.level = lvl
	}
}

// New returns an entry at Debug with room for a few attributes.
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

// Entry is one log record: a level, a message, and slog attributes.
type Entry struct {
	format string
	args   []any
	level  level.Level
	fields []slog.Attr
}

// Reset clears the message, arguments, and attributes and sets the level to Debug.
func (e *Entry) Reset() {
	e.fields = e.fields[:0]
	e.args = e.args[:0]
	e.format = ""
	e.level = level.Debug
}

// Fields returns the attributes. The slice is shared with the entry.
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

// String returns the message and attributes separated by spaces.
func (e *Entry) String() string {
	str := make([]string, len(e.fields)+1)
	str[0] = e.Message()

	for i, field := range e.fields {
		str[i+1] = field.String()
	}

	return strings.Join(str, " ")
}

// Message renders the format and arguments.
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

// Level returns the entry level.
func (e *Entry) Level() level.Level {
	return e.level
}

// SetLevel sets the level.
func (e *Entry) SetLevel(level level.Level) *Entry {
	e.level = level

	return e
}

// SetMessage sets the message and drops format arguments.
func (e *Entry) SetMessage(msg string) *Entry {
	e.format = msg
	e.args = nil

	return e
}

// SetMessagef sets a format and arguments rendered by Message.
func (e *Entry) SetMessagef(format string, args ...any) *Entry {
	e.format = format
	e.args = append(e.args[:0], args...)

	return e
}

// Add appends attributes.
func (e *Entry) Add(fields ...slog.Attr) *Entry {
	e.fields = append(e.fields, fields...)

	return e
}

// AddAny appends an attribute for value.
func (e *Entry) AddAny(key string, value any) *Entry {
	return e.Add(field.Any(key, value))
}

// AddString appends a string attribute.
func (e *Entry) AddString(key, value string) *Entry {
	return e.Add(field.String(key, value))
}

// Replace sets the attribute with key, appending it when the key is missing.
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
