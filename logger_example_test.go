package log

import (
	"bytes"
	"context"
	"log"
)

func ExampleNew() {
	ctx := context.Background()
	buf := &bytes.Buffer{}
	logger := New(NewStdHandler(log.New(buf, "", log.LstdFlags), LevelInfo))
	logger.Info(ctx, "logged info")
}
