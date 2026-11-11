package entry_test

import (
	"log/slog"
	"testing"

	"github.com/go-4devs/log/internal/entry"
	"github.com/go-4devs/log/level"
)

func TestGetResetsReusedEntry(t *testing.T) {
	t.Parallel()

	ent := entry.Get()
	ent.SetLevel(level.Error).SetMessage("kept").Add(slog.String("k", "v"))
	entry.Put(ent)

	again := entry.Get()
	defer entry.Put(again)

	if again != ent {
		t.Fatal("pool returned a new entry")
	}

	if again.Level() != level.Debug || again.Message() != "" || len(again.Fields()) != 0 {
		t.Fatalf("level %v msg %q fields %d", again.Level(), again.Message(), len(again.Fields()))
	}
}
