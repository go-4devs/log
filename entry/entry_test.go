package entry_test

import (
	"testing"
	"time"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
)

func TestEntry_Replace(t *testing.T) {
	t.Parallel()

	ent := entry.New(entry.WithFields(
		field.Any("inti", "init date"),
	))

	ent = ent.Replace("date", field.StringValue("some date"))
	ent = ent.Replace("date", field.TimeValue(time.Time{}))

	fields := ent.Fields()

	if len(fields) != 2 {
		t.Fatalf("count must be 2 got %v", len(fields))
	}

	var has bool

	fields.Fields(func(f field.Field) bool {
		if f.Key == "date" && !f.Value.AsTime().IsZero() {
			has = true

			return false
		}

		return true
	})

	if !has {
		t.Fatal("failed reace value")
	}
}
