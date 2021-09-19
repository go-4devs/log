package log_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"gitoa.ru/go-4devs/log"
)

func TestFields(t *testing.T) {
	type rObj struct {
		id string
	}

	ctx := context.Background()
	buf := &bytes.Buffer{}
	log := log.New(log.WithWriter(buf))
	success := "msg=\"message\" err=file already exists version=0.1.0 obj={id:uid}\n"

	log.InfoKV(ctx, "message",
		"err", os.ErrExist,
		"version", "0.1.0",
		"obj", rObj{id: "uid"},
	)

	if success != buf.String() {
		t.Errorf("invalid value\n got:%s\n exp:%s", buf, success)
	}
}
