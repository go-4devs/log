package logrus_test

import (
	"context"
	"io"
	"os"

	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/handler/logrus"
	slogrus "github.com/sirupsen/logrus"
)

func ExampleNew_logrusHandler() {
	ctx := context.Background()
	lgrs := slogrus.New()
	lgrs.SetOutput(os.Stdout)
	lgrs.SetFormatter(&slogrus.TextFormatter{
		DisableTimestamp: true,
	})

	log := logrus.New(lgrs)
	log.Err(ctx, "log logrus")
	log.ErrKV(ctx, "log logrus kv", field.Int("int", 42))
	log.ErrKVs(ctx, "log logrus kv sugar", "err", io.EOF)

	// Output:
	// level=error msg="log logrus"
	// level=error msg="log logrus kv" int=42
	// level=error msg="log logrus kv sugar" err=EOF
}
