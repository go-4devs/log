package log

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"
	"sync"

	"gitoa.ru/go-4devs/log/entry"
)

// New creates standart logger.
func New(opts ...Option) Logger {
	l := log{e: stringFormat(), w: os.Stderr}

	for _, opt := range opts {
		opt(&l)
	}

	return func(_ context.Context, entry *entry.Entry) (int, error) {
		b, err := l.e(entry)
		if err != nil {
			return 0, err
		}

		return l.w.Write(b)
	}
}

// Option configure log.
type Option func(*log)

// Encode sets formats and encode output message.
type Encode func(*entry.Entry) ([]byte, error)

type log struct {
	w io.Writer
	e Encode
}

// WithWriter sets writer logger.
func WithWriter(writer io.Writer) Option {
	return func(l *log) {
		l.w = writer
	}
}

// WithStdout sets logged to os.Stdout.
func WithStdout() Option {
	return func(l *log) {
		l.w = os.Stdout
	}
}

// WithEncode sets format log.
func WithEncode(e Encode) Option {
	return func(l *log) {
		l.e = e
	}
}

// WithStringFormat sets format as simple string.
func WithStringFormat() Option {
	return func(l *log) {
		l.e = stringFormat()
	}
}

// WithJSONFormat sets json output format.
func WithJSONFormat() Option {
	return func(l *log) {
		l.e = jsonFormat
	}
}

func stringFormat() func(entry *entry.Entry) ([]byte, error) {
	pool := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	return func(entry *entry.Entry) ([]byte, error) {
		b := pool.Get().(*bytes.Buffer)
		b.Reset()

		defer func() {
			pool.Put(b)
		}()

		b.WriteString("msg=\"")
		b.WriteString(strings.TrimSpace(entry.Message()))
		b.WriteString("\"")

		for _, field := range entry.Fields() {
			b.WriteString(" ")
			b.WriteString(string(field.Key()))
			b.WriteString("=")
			b.WriteString(field.Value().String())
		}

		b.WriteString("\n")

		return b.Bytes(), nil
	}
}

func jsonFormat(entry *entry.Entry) ([]byte, error) {
	res, err := json.Marshal(entry.AddString("msg", entry.Message()).Fields().AsMap())
	if err != nil {
		return nil, err
	}

	return append(res, []byte("\n")...), nil
}
