package logrus

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/go-4devs/log"
	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	buf := &bytes.Buffer{}

	lrg := logrus.New()
	lrg.SetOutput(buf)

	handler := New(lrg, log.LevelDebug, WithLevel(log.LevelNotice, logrus.ErrorLevel))
	expect := "level=error msg=\"handle logrus message\"\n"

	handler(ctx, log.LevelNotice, "handle logrus message", log.Fields{})

	if !strings.HasSuffix(buf.String(), expect) {
		t.Errorf("invalid suffix\n got: %s\nexpect:%s\n", buf.String(), expect)
	}
}
