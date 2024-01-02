package log

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
)

func WithSource(depth int) Middleware {
	const offset = 3

	return func(ctx context.Context, data *entry.Entry, handler Logger) (int, error) {
		pc, file, line, has := runtime.Caller(depth + offset)
		if !has {
			return handler(ctx, data.AddAny(KeyLevel, field.NilValue()))
		}

		fnc := runtime.FuncForPC(pc)

		return handler(ctx, data.AddAny(KeySource, Source{
			Func: filepath.Base(fnc.Name()),
			File: filepath.Base(file),
			Line: line,
		}))
	}
}

// Source describes the location of a line of source code.
type Source struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func (l Source) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s:%d", l.File, l.Line)), nil
}

func (l Source) MarshalJSON() ([]byte, error) {
	return fmt.Appendf([]byte{}, `{"file":"%s","line":%d,"func":"%s"}`, l.File, l.Line, l.Func), nil
}
