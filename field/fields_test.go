package field_test

import (
	"encoding/json"
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

func TestFields_Append(t *testing.T) {
	t.Parallel()

	fields := field.Fields{field.Any("any", "value")}
	fields = fields.Append(field.String("string", "value"))

	if len(fields) != 2 {
		t.Fatalf("require 2 field got %v", len(fields))
	}
}

func TestFields_Replace(t *testing.T) {
	t.Parallel()

	fields := field.Fields{
		field.Any("any", "any init"),
		field.Any("replace", "replace init"),
	}

	old, ok := fields.Replace(field.Int64("replace", 42))
	if !ok || old.Key != "replace" || old.Value != field.StringValue("replace init") {
		t.Fatalf("failed replace value:%v", old)
	}

	o2, ok2 := fields.Replace(field.Any("new", "new data"))
	if ok2 || o2.Key != "" {
		t.Fatalf("failed set new data:%v", o2)
	}
}

func TestTimeValueRoundTrip(t *testing.T) {
	t.Parallel()

	if got := field.Time("t", time.Time{}).Value.AsTime(); !got.IsZero() {
		t.Fatalf("zero %s", got)
	}

	far := time.Date(1400, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := field.Time("t", far).Value.AsTime(); !got.Equal(far) {
		t.Fatalf("got %s", got)
	}
}

func TestGroupValueCopiesSlice(t *testing.T) {
	t.Parallel()

	items := []field.Field{field.String("user", "ann")}
	group := field.Groups("g", items...)
	items[0] = field.String("user", "bob")

	if group.Value.AsGroup()[0].Value.AsString() != "ann" {
		t.Fatalf("got %s", group.Value.AsGroup()[0].Value.AsString())
	}
}

func TestAnyBytesAreBinaryAndCopied(t *testing.T) {
	t.Parallel()

	body := []byte("hi")
	val := field.Any("body", body).Value
	body[0] = 'z'

	if val.Kind != field.KindBinary || string(val.AsBinary()) != "hi" {
		t.Fatalf("got %s %q", val.Kind, val.AsBinary())
	}
}

func TestStringsValueCopiesSlice(t *testing.T) {
	t.Parallel()

	names := []string{"ada"}
	val := field.Strings("names", names...).Value
	names[0] = "bob"

	got, _ := val.Any().([]any)
	if len(got) != 1 || got[0] != "ada" {
		t.Fatalf("got %#v", val.Any())
	}
}

func TestKindUnmarshalJSONNumber(t *testing.T) {
	t.Parallel()

	var kind field.Kind

	if err := json.Unmarshal([]byte(`"bytes"`), &kind); err != nil || kind != field.KindBinary {
		t.Fatalf("got %v %v", kind, err)
	}

	if err := json.Unmarshal([]byte("1"), &kind); err == nil {
		t.Fatal("number must return an error")
	}
}
