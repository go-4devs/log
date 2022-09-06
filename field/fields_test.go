package field_test

import (
	"testing"

	"gitoa.ru/go-4devs/log/field"
)

func TestFields_Append(t *testing.T) {
	t.Parallel()

	fields := field.Fields{field.Any("any", "value")}
	fields = fields.Append(field.String("string", "value"))

	if len(fields) != 2 {
		t.Fatalf("require 2 field got %v", len(fields))
	}
}
