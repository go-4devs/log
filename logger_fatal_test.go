package log_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/go-4devs/log"
)

func TestPanicWritesThenPanics(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := log.New(log.WithWriter(buf))

	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()

		logger.Panic("boom")
	}()

	if !bytes.Contains(buf.Bytes(), []byte("boom")) {
		t.Fatalf("panic wrote %q", buf)
	}
}

func TestFatalWritesThenExits(t *testing.T) {
	t.Parallel()

	if os.Getenv("LOG_FATAL") == "1" {
		log.New(log.WithWriter(os.Stdout)).Fatal("down")

		return
	}

	//nolint:gosec // re-exec this test binary
	cmd := exec.CommandContext(context.Background(), os.Args[0], "-test.run=^TestFatalWritesThenExits$")

	cmd.Env = append(os.Environ(), "LOG_FATAL=1")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
		t.Fatalf("exit: %v", err)
	}

	if !bytes.Contains(out.Bytes(), []byte("down")) {
		t.Fatalf("fatal wrote %q", out.Bytes())
	}
}
