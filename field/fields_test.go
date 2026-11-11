package field_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/go-4devs/log/field"
)

func TestNilErrorString(t *testing.T) {
	t.Parallel()

	if got := field.Error("err", nil).Value.String(); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestTimeValueRoundTrip(t *testing.T) {
	t.Parallel()

	if got := field.Time("t", time.Time{}).Value.Time(); !got.IsZero() {
		t.Fatalf("zero %s", got)
	}

	far := time.Date(1400, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := field.Time("t", far).Value.Time(); !got.Equal(far) {
		t.Fatalf("got %s", got)
	}
}

func TestGroupValueCopiesSlice(t *testing.T) {
	t.Parallel()

	items := []slog.Attr{field.String("user", "ann")}
	group := field.Groups("g", items...)
	items[0] = field.String("user", "bob")

	if group.Value.Group()[0].Value.String() != "ann" {
		t.Fatalf("got %s", group.Value.Group()[0].Value.String())
	}
}

func TestAnyBytesAreBinaryAndCopied(t *testing.T) {
	t.Parallel()

	body := []byte("hi")
	val := field.Any("body", body).Value
	body[0] = 'z'

	got, _ := val.Any().([]byte)
	if string(got) != "hi" {
		t.Fatalf("got %s %q", val.Kind(), got)
	}
}

func TestStringsValueCopiesSlice(t *testing.T) {
	t.Parallel()

	names := []string{"ada"}
	val := field.Strings("names", names...).Value
	names[0] = "bob"

	got, _ := field.ValueAny(val).([]any)
	if len(got) != 1 || got[0] != "ada" {
		t.Fatalf("got %#v", val.Any())
	}
}

func TestAsFloat32ReadsFloat64(t *testing.T) {
	t.Parallel()

	if got := field.Float32("x", 1.5).Value.Float64(); got != 1.5 {
		t.Fatalf("got %v", got)
	}
}

func TestAnyIntSliceIsCopied(t *testing.T) {
	t.Parallel()

	nums := []int{1}
	val := field.Any("n", nums).Value
	nums[0] = 2

	got, _ := val.Any().([]int)
	if len(got) != 1 || got[0] != 1 {
		t.Fatalf("got %#v", val.Any())
	}
}
