package log

import (
	"log/slog"

	"github.com/go-4devs/log/field"
)

// Field create field.
func Field(key string, value any) slog.Attr {
	return field.Any(key, value)
}

// FieldError new errors field with key error.
func FieldError(err error) slog.Attr {
	return field.Error("error", err)
}
