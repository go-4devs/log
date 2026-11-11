package field_test

import (
	"math"
	"testing"

	"github.com/go-4devs/log/field"
	"github.com/go-4devs/log/internal/buffer"
)

func TestEncoderJSONAppendField_string(t *testing.T) {
	t.Parallel()

	const expect = `"array":["value","other"],"str":"value","nullableStr":"value","nullstr":null`

	encode := field.NewEncoderJSON()

	buf := buffer.New()

	defer func() {
		buf.Free()
	}()

	val := "value"
	strs := field.Strings("array", val, "other")
	*buf = encode.AppendField(*buf, strs)

	str := field.String("str", val)
	*buf = encode.AppendField(*buf, str)

	strp := field.Stringp("nullableStr", &val)
	*buf = encode.AppendField(*buf, strp)

	nullStr := field.Stringp("nullstr", nil)
	*buf = encode.AppendField(*buf, nullStr)

	if buf.String() != expect {
		t.Errorf("json string expect:%v got:%s", expect, buf)
	}
}

func TestEncoderJSONAppendField_bytesAndNonFinite(t *testing.T) {
	t.Parallel()

	encode := field.NewEncoderJSON()
	got := encode.AppendField(nil, field.Bytes("bin", []byte(`a"b`)))
	got = encode.AppendField(got, field.Float64("nan", math.NaN()))
	got = encode.AppendField(got, field.Float64("pinf", math.Inf(1)))
	got = encode.AppendField(got, field.Float64("ninf", math.Inf(-1)))

	const expect = `"bin":"a\"b","nan":"NaN","pinf":"+Inf","ninf":"-Inf"`
	if string(got) != expect {
		t.Fatalf("got %s", got)
	}
}
