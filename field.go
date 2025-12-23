package log

import (
	"gitoa.ru/go-4devs/log/field"
)

// Field create field.
func Field(key string, value any) field.Field {
	return field.Any(key, value)
}

// FieldError new errors field with key error.
func FieldError(err error) field.Field {
	return field.Error("error", err)
}
