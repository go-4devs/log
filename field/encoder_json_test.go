package field_test

import (
	"testing"

	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/internal/buffer"
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
