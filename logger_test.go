package log

import (
	"bytes"
	"context"
	"log"
	"os"
	"testing"
)

type rObj struct {
	id string
}

func TestFields(t *testing.T) {
	ctx := context.Background()
	buf := &bytes.Buffer{}
	logger := log.New(buf, "", 0)
	log := New(NewStdHandler(logger, LevelDebug))
	success := "level=info msg=\"message\" err=file already exists version=0.1.0 obj={id:uid}\n"

	log.InfoKV(ctx, "message",
		"err", os.ErrExist,
		Field{Key: "version", Value: "0.1.0"},
		"obj", rObj{id: "uid"},
	)

	if success != buf.String() {
		t.Errorf("invalid value\n got:%s\n exp:%s", buf, success)
	}
}
