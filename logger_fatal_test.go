package log //nolint:testpackage // replaces unexported exit

import (
	"bytes"
	"testing"
)

func TestPanicAndFatal(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New(WithWriter(buf))

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

	code := -1
	prev := exit
	exit = func(c int) { code = c }

	t.Cleanup(func() { exit = prev })

	buf.Reset()
	logger.Fatal("down")

	if code != 1 || !bytes.Contains(buf.Bytes(), []byte("down")) {
		t.Fatalf("code %d buf %q", code, buf)
	}
}
