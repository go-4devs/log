// nolint: exhaustruct
package field

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

// StringValue returns a new Value for a string.
func StringValue(value string) Value {
	return Value{
		num:  uint64(len(value)),
		any:  stringptr(unsafe.StringData(value)),
		Kind: KindString,
	}
}

// StringpValue returns a new Value for a *string.
func StringpValue(value *string) Value {
	if value == nil {
		return NilValue()
	}

	return StringValue(*value)
}

// StringpValue returns a new Value for a string.
func StringsValue(value []string) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(value)),
		any: func() []Value {
			values := make([]Value, len(value))
			for idx := range value {
				values[idx] = StringValue(value[idx])
			}

			return values
		},
	}
}

// BoolValue returns a Value for a bool.
func BoolValue(v bool) Value {
	u := uint64(0)
	if v {
		u = 1
	}

	return Value{num: u, Kind: KindBool}
}

// BoolsValue returns a Value for a []bool.
func BoolsValue(values []bool) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = BoolValue(values[idx])
			}

			return vals
		},
	}
}

// BoolpValue returns a new Value for a *bool.
func BoolpValue(value *bool) Value {
	if value == nil {
		return NilValue()
	}

	return BoolValue(*value)
}

// Uint64Value returns a Value for a uint64.
func Uint64Value(v uint64) Value {
	return Value{num: v, Kind: KindUint64}
}

// Uint64sValue returns a Value for a []uint64.
func Uint64sValue(values []uint64) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = Uint64Value(values[idx])
			}

			return vals
		},
	}
}

// Uint8sValue returns a Value for a []uint8.
func Uint8sValue(values []uint8) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = Uint64Value(uint64(values[idx]))
			}

			return vals
		},
	}
}

// Uint64sValue returns a Value for a []uint64.
func Uint64pValue(v *uint64) Value {
	if v == nil {
		return NilValue()
	}

	return Uint64Value(*v)
}

// Int64Value returns a Value for an int64.
func Int64Value(value int64) Value {
	return Value{inum: value, Kind: KindInt64}
}

// Int64sValue returns a Value for an []int64.
func Int64sValue(value []int64) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(value)),
		any: func() []Value {
			vals := make([]Value, len(value))
			for idx := range value {
				vals[idx] = Int64Value(value[idx])
			}

			return vals
		},
	}
}

// Int64sValue returns a Value for an *int64.
func Int64pValue(value *int64) Value {
	if value == nil {
		return NilValue()
	}

	return Int64Value(*value)
}

// Float64Value returns a Value for a floating-point number.
func Float64Value(v float64) Value {
	return Value{num: math.Float64bits(v), Kind: KindFloat64}
}

// Float64Value returns a Value for a floating-points number.
func Float64sValue(values []float64) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = Float64Value(values[idx])
			}

			return vals
		},
	}
}

// Float64Value returns a Value for a floating-points number.
func Float64pValue(v *float64) Value {
	if v == nil {
		return NilValue()
	}

	return Float64Value(*v)
}

// Complex64sValue returns a Value for a []complex64.
func Complex64sValue(values []complex64) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = Complex128Value(complex128(values[idx]))
			}

			return vals
		},
	}
}

// Complex128Value returns a Value for a complex128.
func Complex128Value(v complex128) Value {
	return Value{
		Kind: KindComplex128,
		any:  v,
	}
}

// Complex128Value returns a Value for a []complex128.
func Complex128sValue(values []complex128) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = Complex128Value(values[idx])
			}

			return vals
		},
	}
}

// Complex128Value returns a Value for a *complex128.
func Complex128pValue(v *complex128) Value {
	if v == nil {
		return NilValue()
	}

	return Complex128Value(*v)
}

// TimeValue returns a Value for a time.Time.
func TimeValue(v time.Time) Value {
	return Value{inum: v.UnixNano(), any: v.Location(), Kind: KindTime}
}

// TimepValue returns a Value for a *time.Time.
func TimepValue(v *time.Time) Value {
	if v == nil {
		return NilValue()
	}

	return TimeValue(*v)
}

// TimesValue returns a Value for a []time.Time.
func TimesValue(values []time.Time) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = TimeValue(values[idx])
			}

			return vals
		},
	}
}

func ClosureValue(fn ClosureFn) Value {
	return Value{
		Kind: KindClosure,
		any:  fn,
	}
}

// DurationValue returns a Value for a time.Duration.
func DurationValue(v time.Duration) Value {
	return Value{inum: v.Nanoseconds(), Kind: KindDuration}
}

// DurationValue returns a Value for a *time.Duration.
func DurationpValue(v *time.Duration) Value {
	if v == nil {
		return NilValue()
	}

	return DurationValue(*v)
}

// DurationValue returns a Value for a *time.Duration.
func DurationsValue(values []time.Duration) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(values)),
		any: func() []Value {
			vals := make([]Value, len(values))
			for idx := range values {
				vals[idx] = DurationValue(values[idx])
			}

			return vals
		},
	}
}

// GroupValue returns a new Value for a list of Fields.
func GroupValue(as ...Field) Value {
	return Value{
		num:  uint64(len(as)),
		any:  groupptr(unsafe.SliceData(as)),
		Kind: KindGroup,
	}
}

func ErrorValue(value error) Value {
	return Value{
		Kind: KindError,
		any:  value,
	}
}

func ErrorsValue(value []error) Value {
	return Value{
		Kind: KindArray,
		num:  uint64(len(value)),
		any: func() []Value {
			vals := make([]Value, len(value))
			for idx := range value {
				vals[idx] = ErrorValue(value[idx])
			}

			return vals
		},
	}
}

func BytesValue(value []byte) Value {
	return Value{
		Kind: KindBinary,
		any:  value,
	}
}

//nolint:gochecknoglobals
var nilValue = Value{
	Kind: KindNil,
}

func NilValue() Value {
	return nilValue
}

// AnyValue returns a Value for the supplied value.
//
//nolint:funlen,gocyclo,cyclop
func AnyValue(v any) Value {
	switch value := v.(type) {
	case string:
		return StringValue(value)
	case int:
		return Int64Value(int64(value))
	case uint:
		return Uint64Value(uint64(value))
	case int64:
		return Int64Value(value)
	case *int64:
		return Int64pValue(value)
	case []int64:
		return Int64sValue(value)
	case uint64:
		return Uint64Value(value)
	case *uint64:
		return Uint64pValue(value)
	case []uint64:
		return Uint64sValue(value)
	case bool:
		return BoolValue(value)
	case *bool:
		return BoolpValue(value)
	case []bool:
		return BoolsValue(value)
	case nil:
		return NilValue()
	case complex128:
		return Complex128Value(value)
	case *complex128:
		return Complex128pValue(value)
	case []complex128:
		return Complex128sValue(value)
	case complex64:
		return Complex128Value(complex128(value))
	case []complex64:
		return Complex64sValue(value)
	case time.Duration:
		return DurationValue(value)
	case *time.Duration:
		return DurationpValue(value)
	case []time.Duration:
		return DurationsValue(value)
	case time.Time:
		return TimeValue(value)
	case *time.Time:
		return TimepValue(value)
	case []time.Time:
		return TimesValue(value)
	case uint8:
		return Uint64Value(uint64(value))
	case []uint8:
		return Uint8sValue(value)
	case uint16:
		return Uint64Value(uint64(value))
	case uint32:
		return Uint64Value(uint64(value))
	case uintptr:
		return Uint64Value(uint64(value))
	case int8:
		return Int64Value(int64(value))
	case int16:
		return Int64Value(int64(value))
	case int32:
		return Int64Value(int64(value))
	case float64:
		return Float64Value(value)
	case *float64:
		return Float64pValue(value)
	case []float64:
		return Float64sValue(value)
	case float32:
		return Float64Value(float64(value))
	case error:
		return ErrorValue(value)
	case []error:
		return ErrorsValue(value)
	case []Field:
		return GroupValue(value...)
	case Fields:
		return GroupValue(value...)
	case Kind:
		return Value{Kind: value}
	case func() any:
		return ClosureValue(value)
	case ClosureFn:
		return ClosureValue(value)
	case LogValuer:
		return ClosureValue(value.LogValue)
	case Value:
		return value
	default:
		return Value{any: value}
	}
}

type (
	stringptr *byte  // used in Value.any when the Value is a string
	groupptr  *Field // used in Value.any when the Value is a []Field
)

type Value struct {
	Kind Kind
	num  uint64
	inum int64
	any  any
}

func (v Value) String() string {
	if sp, ok := v.any.(stringptr); ok {
		return unsafe.String(sp, v.num)
	}

	var buf []byte

	return string(v.append(buf))
}

// append appends a text representation of v to dst.
// v is formatted as with fmt.Sprint.
//
//nolint:gomnd,cyclop
func (v Value) append(dst []byte) []byte {
	switch v.Kind {
	case KindString:
		return append(dst, v.AsString()...)
	case KindInt64:
		return strconv.AppendInt(dst, v.inum, 10)
	case KindUint64:
		return strconv.AppendUint(dst, v.num, 10)
	case KindFloat64:
		return strconv.AppendFloat(dst, v.AsFloat64(), 'g', -1, 64)
	case KindFloat32:
		return strconv.AppendFloat(dst, float64(v.AsFloat32()), 'g', -1, 32)
	case KindBool:
		return strconv.AppendBool(dst, v.AsBool())
	case KindDuration:
		return append(dst, v.AsDuration().String()...)
	case KindTime:
		return append(dst, v.AsTime().String()...)
	case KindError:
		return append(dst, v.AsError().Error()...)
	case KindGroup:
		return fmt.Append(dst, v.AsGroup())
	case KindClosure:
		return fmt.Append(dst, v.Resolve())
	case KindAny:
		return fmt.Append(dst, v.any)
	default:
		return fmt.Appendf(dst, "%+v", v.any)
	}
}

// nolint: gocyclo,cyclop
func (v Value) Any() any {
	switch v.Kind {
	case KindAny, KindBinary:
		return v.any
	case KindString:
		return v.AsString()
	case KindInt64:
		return v.AsInt64()
	case KindArray:
		return v.AsArray().Resolve()
	case KindBool:
		return v.AsBool()
	case KindClosure:
		return v.Resolve()
	case KindComplex128:
		return v.AsComplex128()
	case KindDuration:
		return v.AsDuration()
	case KindTime:
		return v.AsTime()
	case KindError:
		return v.AsError()
	case KindFloat32:
		return v.AsFloat32()
	case KindFloat64:
		return v.AsFloat64()
	case KindNil:
		return nil
	case KindUint64:
		return v.AsUint64()
	case KindGroup:
		return v.AsGroup().Any()
	}

	return v.any
}

// nolint: forcetypeassert
func (v Value) AsString() string {
	if v.Kind != KindString {
		return ""
	}

	return unsafe.String(v.any.(stringptr), v.num)
}

func (v Value) AsBool() bool {
	return v.num == 1
}

func (v Value) AsInt64() int64 {
	return v.inum
}

func (v Value) AsUint() uint {
	return uint(v.num)
}

func (v Value) AsUint64() uint64 {
	return v.num
}

func (v Value) AsFloat32() float32 {
	return math.Float32frombits(uint32(v.num))
}

func (v Value) AsFloat64() float64 {
	return math.Float64frombits(v.num)
}

func (v Value) AsComplex128() complex128 {
	cmplex, _ := v.any.(complex128)

	return cmplex
}

func (v Value) AsUintptr() uintptr {
	return uintptr(v.num)
}

func (v Value) AsBinary() []byte {
	bytes, _ := v.any.([]byte)

	return bytes
}

func (v Value) AsDuration() time.Duration {
	return time.Duration(v.inum)
}

func (v Value) AsTime() time.Time {
	loc, ok := v.any.(*time.Location)
	if !ok {
		return time.Time{}
	}

	return time.Unix(0, v.inum).In(loc)
}

func (v Value) AsError() error {
	err, _ := v.any.(error)

	return err
}

//nolint:forcetypeassert
func (v Value) AsGroup() Fields {
	if v.Kind != KindGroup {
		return nil
	}

	return unsafe.Slice((*Field)(v.any.(groupptr)), v.num)
}

func (v Value) Resolve() any {
	cl, ok := v.any.(ClosureFn)
	if !ok {
		return nil
	}

	return cl()
}

type Values []Value

func (v Values) Resolve() any {
	res := make([]any, len(v))
	for idx := range v {
		res[idx] = v[idx].Any()
	}

	return res
}

func (v Value) AsArray() Values {
	switch res := v.any.(type) {
	case []Value:
		return res
	case func() []Value:
		return res()
	default:
		return nil
	}
}
