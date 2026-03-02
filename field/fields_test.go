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

func TestFields_Replace(t *testing.T) {
	t.Parallel()

	fields := field.Fields{
		field.Any("any", "any init"),
		field.Any("replace", "replace init"),
	}

	old, ok := fields.Replace(field.Int64("replace", 42))
	if !ok || old.Key != "replace" || old.Value != field.StringValue("replace init") {
		t.Fatalf("failed replace value:%v", old)
	}

	o2, ok2 := fields.Replace(field.Any("new", "new data"))
	if ok2 || o2.Key != "" {
		t.Fatalf("failed set new data:%v", o2)
	}
}
