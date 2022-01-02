package log_test

import (
	"io"
	"os"

	slogrus "github.com/sirupsen/logrus"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/handler/logrus"
)

func ExampleNew_logrusHandler() {
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
