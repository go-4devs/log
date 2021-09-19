package log

import (
	"fmt"
	"strings"
)

// Fields slice field.
type Fields []Field

// String implement stringer.
func (f Fields) String() string {
	str := make([]string, len(f))
	for i, field := range f {
		str[i] = field.String()
	}

	return strings.Join(str, " ")
}

// NewField create field.
func NewField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// Field struct.
type Field struct {
	Key   string
	Value interface{}
}

// String implent stringer.
func (f Field) String() string {
	return fmt.Sprintf("%s=%+v", f.Key, f.Value)
}

// FieldError new errors field with key error.
func FieldError(err error) Field {
	return NewField("error", err)
}
