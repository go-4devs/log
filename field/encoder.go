package field

import (
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	base10 = 10

	float64Bits    = 64
	complex128Bits = 128
)

type Encoder interface {
	AppendField(dst []byte, attr slog.Attr) []byte
	AppendValue(dst []byte, val slog.Value) []byte
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

// WithGropuConfig set group config.
//
// Deprecated: use WithGroupConfig.
func WithGropuConfig(start, end, deli byte) func(*BaseEncoder) {
	return WithGroupConfig(start, end, deli)
}

func WithGroupConfig(start, end, deli byte) func(*BaseEncoder) {
	return func(be *BaseEncoder) {
		be.group = groupConfig{
			start: start,
			end:   end,
			deli:  deli,
		}
	}
}

func WithDefaultValue(fn func(dst []byte, e Encoder, val slog.Value) []byte) func(*BaseEncoder) {
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
		DefaultValue: func(dst []byte, e Encoder, val slog.Value) []byte {
			return e.AppendValue(dst, slog.StringValue(fmt.Sprintf("%+v", val.Any())))
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
	DefaultValue func(dst []byte, e Encoder, val slog.Value) []byte
}

func (b BaseEncoder) AppendValue(dst []byte, val slog.Value) []byte {
	return b.appendValue(dst, val, "", 0)
}

func (b BaseEncoder) AppendDelimiter(dst []byte, deli byte) []byte {
	if deli == 0 {
		return dst
	}

	return append(dst, deli)
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

func (b BaseEncoder) AppendField(dst []byte, attr slog.Attr) []byte {
	prefix := ""

	if len(dst) != 0 {
		prew := dst[len(dst)-1]
		if prew != '{' && prew != '.' {
			prefix = string(b.group.deli)
		}
	}

	return b.appendField(dst, attr, prefix, b.delimeter)
}

func (b BaseEncoder) AppendKey(dst []byte, key string, prefix string) []byte {
	// A bare group delimiter stays outside the key. A dotted path is part of the key,
	// so JSON quotes "user.name" instead of writing user."name".
	if prefix != "" && prefix != string(b.group.deli) {
		sep, path := cutSep(prefix, b.group.deli)
		if sep != "" {
			dst = append(dst, sep...)
		}

		return b.AppendString(dst, path+key)
	}

	if prefix != "" {
		dst = append(dst, prefix...)
	}

	return b.AppendString(dst, key)
}

func (b BaseEncoder) AppendComplex(dst []byte, c complex128) []byte {
	cmplx := strconv.FormatComplex(c, 'g', -1, complex128Bits)

	return b.AppendString(dst, cmplx)
}

func (b BaseEncoder) AppendFloat(dst []byte, f float64, bitSize int) []byte {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return b.AppendString(dst, strconv.FormatFloat(f, 'g', -1, bitSize))
	}

	return strconv.AppendFloat(dst, f, 'g', -1, bitSize)
}

func (b BaseEncoder) AppendUint(dst []byte, u uint64) []byte {
	return strconv.AppendUint(dst, u, base10)
}

func (b BaseEncoder) AppendNull(dst []byte) []byte {
	return append(dst, b.nullValue...)
}

func (b BaseEncoder) AppendInt(dst []byte, val int64) []byte {
	return strconv.AppendInt(dst, val, base10)
}

func (b BaseEncoder) AppendBool(dst []byte, val bool) []byte {
	return strconv.AppendBool(dst, val)
}

func (b BaseEncoder) AppendGroup(dst []byte, fields []slog.Attr) []byte {
	dst = append(dst, b.group.start)
	dst = b.appendGroup(dst, fields, "")

	return append(dst, b.group.end)
}

func (b BaseEncoder) AppendArray(dst []byte, in []slog.Value) []byte {
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
	return strconv.AppendQuote(dst, string(in))
}

func (b BaseEncoder) appendGroup(dst []byte, fields []slog.Attr, path string) []byte {
	sep, rest := cutSep(path, b.group.deli)

	for i, field := range fields {
		next := rest
		if i == 0 {
			next = sep + rest
		} else if b.group.deli != 0 {
			next = string(b.group.deli) + rest
		}

		dst = b.appendField(dst, field, next, b.delimeter)
	}

	return dst
}

func (b BaseEncoder) appendField(dst []byte, attr slog.Attr, prefix string, deli byte) []byte {
	val := attr.Value.Resolve()
	if val.Kind() == slog.KindGroup {
		return b.appendGroup(dst, val.Group(), prefix+attr.Key+".")
	}

	dst = b.AppendKey(dst, attr.Key, prefix)

	return b.appendValue(dst, val, attr.Key+".", deli)
}

func cutSep(prefix string, deli byte) (string, string) {
	if deli != 0 && prefix != "" && prefix[0] == deli {
		return prefix[:1], prefix[1:]
	}

	return "", prefix
}

//nolint:cyclop
func (b BaseEncoder) appendValue(dst []byte, val slog.Value, prefix string, deli byte) []byte {
	val = val.Resolve()

	switch val.Kind() {
	case slog.KindGroup:
		return b.appendGroup(dst, val.Group(), prefix)
	case slog.KindBool:
		return b.AppendBool(b.AppendDelimiter(dst, deli), val.Bool())
	case slog.KindInt64:
		return b.AppendInt(b.AppendDelimiter(dst, deli), val.Int64())
	case slog.KindFloat64:
		return b.AppendFloat(b.AppendDelimiter(dst, deli), val.Float64(), float64Bits)
	case slog.KindUint64:
		return b.AppendUint(b.AppendDelimiter(dst, deli), val.Uint64())
	case slog.KindString:
		return b.AppendString(b.AppendDelimiter(dst, deli), val.String())
	case slog.KindDuration:
		return b.AppendDuration(b.AppendDelimiter(dst, deli), val.Duration())
	case slog.KindTime:
		return b.AppendTime(b.AppendDelimiter(dst, deli), val.Time())
	case slog.KindAny:
		return b.appendAny(dst, val, deli)
	default:
		return b.DefaultValue(b.AppendDelimiter(dst, deli), b, val)
	}
}

func (b BaseEncoder) appendAny(dst []byte, val slog.Value, deli byte) []byte {
	switch value := val.Any().(type) {
	case nil:
		return b.AppendNull(b.AppendDelimiter(dst, deli))
	case []byte:
		return b.AppendBytes(b.AppendDelimiter(dst, deli), value)
	case list:
		return b.AppendArray(b.AppendDelimiter(dst, deli), value)
	case error:
		text := ""
		if value != nil {
			text = value.Error()
		}

		return b.AppendString(b.AppendDelimiter(dst, deli), text)
	case complex128:
		return b.AppendComplex(b.AppendDelimiter(dst, deli), value)
	case complex64:
		return b.AppendComplex(b.AppendDelimiter(dst, deli), complex128(value))
	default:
		return b.DefaultValue(b.AppendDelimiter(dst, deli), b, val)
	}
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
