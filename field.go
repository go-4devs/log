package log

import (
	"log/slog"

	"github.com/go-4devs/log/field"
)

// Field returns an slog attribute for value.
func Field(key string, value any) slog.Attr {
	return field.Any(key, value)
}

// FieldError returns an error attribute with the key "error".
func FieldError(err error) slog.Attr {
	return field.Error("error", err)
}
