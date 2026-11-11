package log

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
)

const funcName = "Logger"

// WithSource adds the first caller outside this module as a source attribute.
func WithSource(items int, trimPath func(string) string) Middleware {
	const (
		skip    = 4
		pkgName = "github.com/go-4devs/log"
	)

	return WithCallers(items, skip, pkgName, trimPath)
}

// WithCallers adds the first caller outside logPkg as a source attribute.
func WithCallers(items, skip int, logPkg string, trimPath func(string) string) Middleware {
	items += skip
	skipHelper := logPkg + "."
	funcPrefix := skipHelper + funcName

	return func(ctx context.Context, data *entry.Entry, handler Logger) (int, error) {
		pc := make([]uintptr, items)
		n := runtime.Callers(skip, pc)

		if n == 0 {
			return handler(ctx, data.Add(errSourceField(skip, items)))
		}

		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
		frames := runtime.CallersFrames(pc)
		prew := false

		for {
			frame, more := frames.Next()

			has := strings.HasPrefix(frame.Function, funcPrefix)
			if !has && prew {
				if strings.HasPrefix(frame.Function, skipHelper) {
					continue
				}

				return handler(ctx, data.AddAny(KeySource, Source{
					Func: filepath.Base(frame.Function),
					Line: frame.Line,
					File: trimPath(frame.File),
				}))
			}

			prew = has

			if !more {
				break
			}
		}

		return handler(ctx, data.Add(errSourceField(skip, items)))
	}
}

// TrimPath returns the last two path elements of file.
func TrimPath(file string) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return filepath.Base(file)
	}

	// Find the penultimate separator.
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return filepath.Base(file)
	}

	return file[idx+1:]
}

// Source describes the location of a line of source code.
type Source struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}

// MarshalText encodes the source as file:line.
func (l Source) MarshalText() ([]byte, error) {
	return fmt.Appendf(nil, "%s:%d", l.File, l.Line), nil
}

// MarshalJSON encodes the file, line, and function.
func (l Source) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `{"file":%q,"line":%d,"func":%q}`, l.File, l.Line, l.Func), nil
}

func errSourceField(skip, mframe int) slog.Attr {
	return field.String(KeySource, fmt.Sprintf("source not found by frames[%d:%d]", skip, mframe))
}
