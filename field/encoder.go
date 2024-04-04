//nolint:gomnd
package field

import (
	"fmt"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

type Encoder interface {
	AppendField(dst []byte, field Field) []byte
	AppendValue(dst []byte, val Value) []byte
}

func WithAppendString(fn func(dst []byte, in string) []byte) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.AppendString = fn
	}
}

func WithNullValue(in string) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.nullValue = []byte(in)
	}
}

func WithDelimeter(in byte) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.delimeter = in
	}
}

func WithGropuConfig(start, end, deli byte) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.group = groupConfig{
			start: start,
			end:   end,
			deli:  deli,
		}
	}
}

func WithDefaultValue(fn func(dst []byte, e Encoder, val Value) []byte) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.DefaultValue = fn
	}
}

func NewEncoder(opts ...func(*BaseEncoder)) BaseEncoder {
	be := BaseEncoder{
		nullValue: []byte("null"),
		group: groupConfig{
			start: '{',
			end:   '}',
			deli:  ',',
		},
		array: groupConfig{
			start: '[',
			end:   ']',
			deli:  ',',
		},
		timeFormat:   time.RFC3339,
		AppendString: AppendString,
		delimeter:    '=',
		DefaultValue: func(dst []byte, e Encoder, val Value) []byte {
			return e.AppendValue(dst, StringValue(fmt.Sprintf("%+v", val.Any())))
		},
	}

	for _, opt := range opts {
		opt(&be)
	}

	return be
}

type groupConfig struct {
	start byte
	end   byte
	deli  byte
}

type BaseEncoder struct {
	nullValue    []byte
	group        groupConfig
	array        groupConfig
	timeFormat   string
	AppendString func(dst []byte, in string) []byte
	delimeter    byte
	DefaultValue func(dst []byte, e Encoder, val Value) []byte
}

func (b BaseEncoder) AppendValue(dst []byte, val Value) []byte {
	return b.appendValue(dst, val, "", 0)
}

func (b BaseEncoder) AppendDelimiter(dst []byte, deli byte) []byte {
	if deli == 0 {
		return dst
	}

	return append(dst, deli)
}

//nolint:gocyclo,cyclop
func (b BaseEncoder) appendValue(dst []byte, val Value, prefix string, deli byte) []byte {
	switch val.Kind {
	case KindGroup:
		return b.appendGroup(dst, val.AsGroup(), prefix)
	case KindClosure:
		return b.appendValue(dst, AnyValue(val.Resolve()), prefix, deli)
	case KindArray:
		return b.AppendArray(b.AppendDelimiter(dst, deli), val.AsArray())
	case KindNil:
		return b.AppendNull(b.AppendDelimiter(dst, deli))
	case KindBool:
		return b.AppendBool(b.AppendDelimiter(dst, deli), val.AsBool())
	case KindBinary:
		return b.AppendBytes(b.AppendDelimiter(dst, deli), val.AsBinary())
	case KindComplex128:
		return b.AppendComplex(b.AppendDelimiter(dst, deli), val.AsComplex128())
	case KindInt64:
		return b.AppendInt(b.AppendDelimiter(dst, deli), val.AsInt64())
	case KindFloat32:
		return b.AppendFloat(b.AppendDelimiter(dst, deli), float64(val.AsFloat32()), 32)
	case KindFloat64:
		return b.AppendFloat(b.AppendDelimiter(dst, deli), val.AsFloat64(), 64)
	case KindUint64:
		return b.AppendUint(b.AppendDelimiter(dst, deli), val.AsUint64())
	case KindError:
		return b.AppendString(b.AppendDelimiter(dst, deli), val.AsError().Error())
	case KindString:
		return b.AppendString(b.AppendDelimiter(dst, deli), val.AsString())
	case KindDuration:
		return b.AppendDuration(b.AppendDelimiter(dst, deli), val.AsDuration())
	case KindTime:
		return b.AppendTime(b.AppendDelimiter(dst, deli), val.AsTime())
	case KindAny:
		return b.DefaultValue(b.AppendDelimiter(dst, deli), b, val)
	}

	return b.DefaultValue(b.AppendDelimiter(dst, deli), b, val)
}

func (b BaseEncoder) AppendDuration(dst []byte, d time.Duration) []byte {
	return b.AppendString(dst, d.String())
}

func (b BaseEncoder) AppendTime(dst []byte, t time.Time) []byte {
	return b.AppendString(dst, t.Format(b.timeFormat))
}

func AppendString(dst []byte, in string) []byte {
	if needsQuoting(in) {
		return strconv.AppendQuote(dst, in)
	}

	return append(dst, in...)
}

//nolint:cyclop
func needsQuoting(in string) bool {
	if len(in) == 0 {
		return true
	}

	for i := 0; i < len(in); {
		char := in[i]
		if char < utf8.RuneSelf {
			// Quote anything except a backslash that would need quoting in a
			// JSON string, as well as space and '='
			if char != '\\' && (char == ' ' || char == '=' || !safeSet[char]) {
				return true
			}

			i++

			continue
		}

		decodeRune, size := utf8.DecodeRuneInString(in[i:])
		if decodeRune == utf8.RuneError || unicode.IsSpace(decodeRune) || !unicode.IsPrint(decodeRune) {
			return true
		}

		i += size
	}

	return false
}

func (b BaseEncoder) AppendField(dst []byte, field Field) []byte {
	prefix := ""

	if len(dst) != 0 {
		prew := dst[len(dst)-1]
		if prew != '{' && prew != '.' {
			prefix = string(b.group.deli)
		}
	}

	return b.appendField(dst, field, prefix, b.delimeter)
}

func (b BaseEncoder) appendField(dst []byte, field Field, prefix string, deli byte) []byte {
	dst = b.AppendKey(dst, field.Key, prefix)

	return b.appendValue(dst, field.Value, field.Key+".", deli)
}

func (b BaseEncoder) AppendKey(dst []byte, key string, prefix string) []byte {
	if prefix != "" {
		dst = append(dst, prefix...)
	}

	return b.AppendString(dst, key)
}

func (b BaseEncoder) AppendComplex(dst []byte, c complex128) []byte {
	cmplx := strconv.FormatComplex(c, 'g', -1, 128)

	return b.AppendString(dst, cmplx)
}

func (b BaseEncoder) AppendFloat(dst []byte, f float64, bitSize int) []byte {
	return strconv.AppendFloat(dst, f, 'g', -1, bitSize)
}

func (b BaseEncoder) AppendUint(dst []byte, u uint64) []byte {
	return strconv.AppendUint(dst, u, 10)
}

func (b BaseEncoder) AppendNull(dst []byte) []byte {
	return append(dst, b.nullValue...)
}

func (b BaseEncoder) AppendInt(dst []byte, val int64) []byte {
	return strconv.AppendInt(dst, val, 10)
}

func (b BaseEncoder) AppendBool(dst []byte, val bool) []byte {
	return strconv.AppendBool(dst, val)
}

func (b BaseEncoder) AppendGroup(dst []byte, fields []Field) []byte {
	dst = append(dst, b.group.start)
	dst = b.appendGroup(dst, fields, "")

	return append(dst, b.group.end)
}

func (b BaseEncoder) appendGroup(dst []byte, fields []Field, prefix string) []byte {
	if len(fields) > 0 {
		dst = b.appendField(dst, fields[0], ".", b.delimeter)
		for _, field := range fields[1:] {
			dst = b.appendField(append(dst, b.group.deli), field, prefix, b.delimeter)
		}
	}

	return dst
}

func (b BaseEncoder) AppendArray(dst []byte, in []Value) []byte {
	dst = append(dst, b.array.start)
	if len(in) > 0 {
		dst = b.appendValue(dst, in[0], "", 0)
		for _, value := range in[1:] {
			dst = b.appendValue(append(dst, b.array.deli), value, "", 0)
		}
	}

	return append(dst, b.array.end)
}

func (b BaseEncoder) AppendBytes(dst, in []byte) []byte {
	dst = append(dst, '"')
	dst = append(dst, in...)

	return append(dst, '"')
}
