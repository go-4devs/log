package otel

import (
	"testing"

	"github.com/go-4devs/log/level"
)

func TestLevelsOrder(t *testing.T) {
	t.Parallel()

	if levels(level.Emergency) != levelFatal {
		t.Fatalf("emergency is %v", levels(level.Emergency))
	}

	if levels(level.Alert) != levelError3 {
		t.Fatalf("alert is %v", levels(level.Alert))
	}
}
