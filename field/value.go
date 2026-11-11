package field

import (
	"log/slog"
	"reflect"
	"time"
)

// list is a JSON/text array. slog has no array kind.
type list []slog.Value

// StringpValue returns the string, or a nil value when value is nil.
func StringpValue(value *string) slog.Value {
	if value == nil {
		return slog.AnyValue(nil)
	}

	return slog.StringValue(*value)
}

// StringsValue returns a copied string array.
func StringsValue(value []string) slog.Value {
	return valuesOf(value, slog.StringValue)
}

// BoolsValue returns a copied bool array.
func BoolsValue(values []bool) slog.Value {
	return valuesOf(values, slog.BoolValue)
}

// BoolpValue returns the bool, or a nil value when value is nil.
func BoolpValue(value *bool) slog.Value {
	if value == nil {
		return slog.AnyValue(nil)
	}

	return slog.BoolValue(*value)
}

// Uint64sValue returns a copied uint64 array.
func Uint64sValue(values []uint64) slog.Value {
	return valuesOf(values, slog.Uint64Value)
}

// Uint8sValue returns a copied array of numbers, not a byte string.
func Uint8sValue(values []uint8) slog.Value {
	return valuesOf(values, func(v uint8) slog.Value {
		return slog.Uint64Value(uint64(v))
	})
}

// Uint64pValue returns the uint64, or a nil value when v is nil.
func Uint64pValue(v *uint64) slog.Value {
	if v == nil {
		return slog.AnyValue(nil)
	}

	return slog.Uint64Value(*v)
}

// Int64sValue returns a copied int64 array.
func Int64sValue(value []int64) slog.Value {
	return valuesOf(value, slog.Int64Value)
}

// Int64pValue returns the int64, or a nil value when value is nil.
func Int64pValue(value *int64) slog.Value {
	if value == nil {
		return slog.AnyValue(nil)
	}

	return slog.Int64Value(*value)
}

// Float64sValue returns a copied float64 array.
func Float64sValue(values []float64) slog.Value {
	return valuesOf(values, slog.Float64Value)
}

// Float64pValue returns the float64, or a nil value when v is nil.
func Float64pValue(v *float64) slog.Value {
	if v == nil {
		return slog.AnyValue(nil)
	}

	return slog.Float64Value(*v)
}

// Complex64sValue returns a copied complex64 array.
func Complex64sValue(values []complex64) slog.Value {
	return valuesOf(values, func(v complex64) slog.Value {
		return slog.AnyValue(complex128(v))
	})
}

// Complex128sValue returns a copied complex128 array.
func Complex128sValue(values []complex128) slog.Value {
	return valuesOf(values, func(v complex128) slog.Value {
		return slog.AnyValue(v)
	})
}

// Complex128pValue returns the complex128, or a nil value when v is nil.
func Complex128pValue(v *complex128) slog.Value {
	if v == nil {
		return slog.AnyValue(nil)
	}

	return slog.AnyValue(*v)
}

// TimepValue returns the time, or a nil value when v is nil.
func TimepValue(v *time.Time) slog.Value {
	if v == nil {
		return slog.AnyValue(nil)
	}

	return slog.TimeValue(*v)
}

// TimesValue returns a copied time array.
func TimesValue(values []time.Time) slog.Value {
	return valuesOf(values, slog.TimeValue)
}

// DurationpValue returns the duration, or a nil value when v is nil.
func DurationpValue(v *time.Duration) slog.Value {
	if v == nil {
		return slog.AnyValue(nil)
	}

	return slog.DurationValue(*v)
}

// DurationsValue returns a copied duration array.
func DurationsValue(values []time.Duration) slog.Value {
	return valuesOf(values, slog.DurationValue)
}

// GroupValue returns a group. The attribute slice is copied.
func GroupValue(as ...slog.Attr) slog.Value {
	as = append([]slog.Attr(nil), as...)

	return slog.GroupValue(as...)
}

// ErrorValue returns the error text. A nil error is an empty string.
func ErrorValue(value error) slog.Value {
	if value == nil {
		return slog.StringValue("")
	}

	return slog.AnyValue(value)
}

// ErrorsValue returns a copied array of errors.
func ErrorsValue(value []error) slog.Value {
	return valuesOf(value, ErrorValue)
}

// BytesValue returns a copy of value stored as bytes.
func BytesValue(value []byte) slog.Value {
	return slog.AnyValue(copied(value))
}

// ClosureValue returns a value resolved by calling fn when logged.
func ClosureValue(fn ClosureFn) slog.Value {
	return slog.AnyValue(fn)
}

// AnyValue returns a slog value.
// Pointers to nil become a nil value. Slices are copied.
//
//nolint:funlen,gocyclo,cyclop
func AnyValue(v any) slog.Value {
	switch value := v.(type) {
	case string:
		return slog.StringValue(value)
	case int:
		return slog.Int64Value(int64(value))
	case uint:
		return slog.Uint64Value(uint64(value))
	case int64:
		return slog.Int64Value(value)
	case *int64:
		return Int64pValue(value)
	case []int64:
		return Int64sValue(value)
	case uint64:
		return slog.Uint64Value(value)
	case *uint64:
		return Uint64pValue(value)
	case []uint64:
		return Uint64sValue(value)
	case bool:
		return slog.BoolValue(value)
	case *bool:
		return BoolpValue(value)
	case []bool:
		return BoolsValue(value)
	case nil:
		return slog.AnyValue(nil)
	case complex128:
		return slog.AnyValue(value)
	case *complex128:
		return Complex128pValue(value)
	case []complex128:
		return Complex128sValue(value)
	case complex64:
		return slog.AnyValue(complex128(value))
	case []complex64:
		return Complex64sValue(value)
	case time.Duration:
		return slog.DurationValue(value)
	case *time.Duration:
		return DurationpValue(value)
	case []time.Duration:
		return DurationsValue(value)
	case time.Time:
		return slog.TimeValue(value)
	case *time.Time:
		return TimepValue(value)
	case []time.Time:
		return TimesValue(value)
	case uint8:
		return slog.Uint64Value(uint64(value))
	case []uint8:
		return BytesValue(value)
	case uint16:
		return slog.Uint64Value(uint64(value))
	case uint32:
		return slog.Uint64Value(uint64(value))
	case uintptr:
		return slog.Uint64Value(uint64(value))
	case int8:
		return slog.Int64Value(int64(value))
	case int16:
		return slog.Int64Value(int64(value))
	case int32:
		return slog.Int64Value(int64(value))
	case float64:
		return slog.Float64Value(value)
	case *float64:
		return Float64pValue(value)
	case []float64:
		return Float64sValue(value)
	case float32:
		return slog.Float64Value(float64(value))
	case error:
		return ErrorValue(value)
	case []error:
		return ErrorsValue(value)
	case []slog.Attr:
		return GroupValue(value...)
	case func() any:
		return ClosureValue(value)
	case ClosureFn:
		return ClosureValue(value)
	case slog.LogValuer:
		return slog.AnyValue(value)
	case slog.Value:
		return value
	default:
		return slog.AnyValue(detachSlice(value))
	}
}

// ValueAny resolves val. Groups become maps and arrays become slices.
func ValueAny(val slog.Value) any {
	val = val.Resolve()
	if val.Kind() == slog.KindGroup {
		group := val.Group()
		fields := make(map[string]any, len(group))

		for _, attr := range group {
			fields[attr.Key] = ValueAny(attr.Value)
		}

		return fields
	}

	if vals, ok := val.Any().(list); ok {
		out := make([]any, len(vals))
		for i := range vals {
			out[i] = ValueAny(vals[i])
		}

		return out
	}

	return val.Any()
}

func valuesOf[T any](in []T, conv func(T) slog.Value) slog.Value {
	in = copied(in)
	vals := make(list, len(in))

	for i := range in {
		vals[i] = conv(in[i])
	}

	return slog.AnyValue(vals)
}

func copied[T any](in []T) []T {
	return append([]T(nil), in...)
}

func detachSlice(value any) any {
	src := reflect.ValueOf(value)
	if src.Kind() != reflect.Slice || src.IsNil() {
		return value
	}

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	return dst.Interface()
}
