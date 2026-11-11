package logrus_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/go-4devs/log/entry"
	"github.com/go-4devs/log/handler/logrus"
	"github.com/go-4devs/log/level"
	lgr "github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := &bytes.Buffer{}

	lgrus := lgr.New()
	lgrus.SetLevel(lgr.DebugLevel)
	lgrus.SetOutput(buf)
	lgrus.SetFormatter(&lgr.TextFormatter{
		DisableTimestamp: true,
	})

	handler := logrus.New(lgrus)
	expect := "level=info msg=\"handle logrus message\"\n"

	if _, err := handler(ctx, entry.New(entry.WithLevel(level.Info), entry.WithMessage("handle logrus message"))); err != nil {
		t.Error(err)
	}

	if !strings.HasSuffix(buf.String(), expect) {
		t.Errorf("invalid suffix\n got: %s\nexpect:%s\n", buf.String(), expect)
	}
}
