package log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"gitoa.ru/go-4devs/log/entry"
)

// New creates standart logger.
func New(opts ...Option) Logger {
	logger := log{e: stringFormat(), w: os.Stderr}

	for _, opt := range opts {
		opt(&logger)
	}

	return func(_ context.Context, entry *entry.Entry) (int, error) {
		b, err := logger.e(entry)
		if err != nil {
			return 0, fmt.Errorf("enode err: %w", err)
		}

		n, err := logger.w.Write(b)
		if err != nil {
			return 0, fmt.Errorf("failed write: %w", err)
		}

		return n, nil
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
	return WithWriter(os.Stdout)
}

// WithEncode sets format log.
func WithEncode(e Encode) Option {
	return func(l *log) {
		l.e = e
	}
}

// WithStringFormat sets format as simple string.
func WithStringFormat() Option {
	return WithEncode(stringFormat())
}

// WithJSONFormat sets json output format.
func WithJSONFormat() Option {
	return WithEncode(jsonFormat)
}

//nolint: forcetypeassert
func stringFormat() func(entry *entry.Entry) ([]byte, error) {
	pool := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	return func(entry *entry.Entry) ([]byte, error) {
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset()

		defer func() {
			pool.Put(buf)
		}()

		buf.WriteString("msg=\"")
		buf.WriteString(strings.TrimSpace(entry.Message()))
		buf.WriteString("\"")

		for _, field := range entry.Fields() {
			buf.WriteString(" ")
			buf.WriteString(string(field.Key()))
			buf.WriteString("=")
			buf.WriteString(field.Value().String())
		}

		buf.WriteString("\n")

		return buf.Bytes(), nil
	}
}

func jsonFormat(entry *entry.Entry) ([]byte, error) {
	res, err := json.Marshal(entry.AddString("msg", entry.Message()).Fields().AsMap())
	if err != nil {
		return nil, fmt.Errorf("marshal err: %w", err)
	}

	return append(res, []byte("\n")...), nil
}
