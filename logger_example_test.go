package log

import (
	"bytes"
	"context"
)

func ExampleNew() {
	ctx := context.Background()
	buf := &bytes.Buffer{}
	logger := New(NewStdHandler(buf, WithStdLevel(LevelInfo)))
	logger.Info(ctx, "logged info")
}
