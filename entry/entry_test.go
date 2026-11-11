package entry_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/level"
)

func TestEntry_Replace(t *testing.T) {
	t.Parallel()

	ent := entry.New(entry.WithFields(
		field.Any("inti", "init date"),
	))

	ent = ent.Replace("date", slog.StringValue("some date"))
	ent = ent.Replace("date", slog.TimeValue(time.Time{}))

	fields := ent.Fields()

	if len(fields) != 2 {
		t.Fatalf("count must be 2 got %v", len(fields))
	}

	var has bool

	for _, attr := range fields {
		if attr.Key == "date" && attr.Value.Kind() == slog.KindTime && attr.Value.Time().IsZero() {
			has = true

			break
		}
	}

	if !has {
		t.Fatal("failed reace value")
	}
}

func TestSetMessageDropsArgs(t *testing.T) {
	t.Parallel()

	ent := entry.New()
	ent.SetMessagef("id=%d", 7)
	ent.SetMessage("ready")

	if ent.Message() != "ready" {
		t.Fatalf("got %q", ent.Message())
	}

	ent = entry.New(entry.WithMessagef("id=%d", 7), entry.WithMessage("ready"))
	if ent.Message() != "ready" {
		t.Fatalf("option got %q", ent.Message())
	}
}

func TestEntryResetLevel(t *testing.T) {
	t.Parallel()

	ent := entry.New()
	ent.SetLevel(level.Error)
	ent.Reset()

	if ent.Level() != level.Debug {
		t.Fatalf("got %s", ent.Level())
	}
}
