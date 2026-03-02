package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
)

const funcName = "Logger"

func WithSource(items int, trimPath func(string) string) Middleware {
	const (
		skip    = 4
		pkgName = "gitoa.ru/go-4devs/log"
	)

	return WithCallers(items, skip, pkgName, trimPath)
}

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

func (l Source) MarshalText() ([]byte, error) {
	return fmt.Appendf(nil, "%s:%d", l.File, l.Line), nil
}

func (l Source) MarshalJSON() ([]byte, error) {
	return fmt.Appendf([]byte{}, `{"file":%q,"line":%d,"func":"%s"}`, l.File, l.Line, l.Func), nil
}

func errSourceField(skip, mframe int) field.Field {
	return field.String(KeySource, fmt.Sprintf("source not found by frames[%d:%d]", skip, mframe))
}
