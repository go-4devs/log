// Package field builds slog attributes.
// Slice helpers copy their input. Stringp, Boolp, Int64p, Uint64p, Float64p,
// Complex128p, Timep, and Durationp turn a nil pointer into a nil value.
package field

import (
	"log/slog"
	"time"
)

// Empty returns a zero attribute.
func Empty() slog.Attr {
	var attr slog.Attr

	return attr
}

// Any returns an attribute. Nil pointers of known types become a nil value and slices are copied.
func Any(key string, value any) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// String returns an attribute.
func String(key, value string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.StringValue(value),
	}
}

// Stringp returns an attribute. A nil pointer is a nil value.
func Stringp(key string, value *string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: StringpValue(value),
	}
}

// Strings returns an array attribute. The slice is copied.
func Strings(key string, value ...string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: StringsValue(value),
	}
}

// Bool returns an attribute.
func Bool(key string, value bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.BoolValue(value),
	}
}

// Bools returns an array attribute. The slice is copied.
func Bools(key string, value ...bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BoolsValue(value),
	}
}

// Boolp returns an attribute. A nil pointer is a nil value.
func Boolp(key string, value *bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BoolpValue(value),
	}
}

// Uint returns an attribute.
func Uint(key string, value uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

// Uints returns an array attribute. The slice is copied.
func Uints(key string, value ...uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uintp returns an attribute for the pointed value.
func Uintp(key string, value *uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint8 returns an attribute.
func Uint8(key string, value uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

// Uint8s returns an array attribute. The slice is copied.
func Uint8s(key string, value ...uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint8sValue(value),
	}
}

// Uint8p returns an attribute for the pointed value.
func Uint8p(key string, value *uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint16 returns an attribute.
func Uint16(key string, value uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

// Uint16s returns an array attribute. The slice is copied.
func Uint16s(key string, value ...uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint16p returns an attribute for the pointed value.
func Uint16p(key string, value *uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint32 returns an attribute.
func Uint32(key string, value uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

// Uint32s returns an array attribute. The slice is copied.
func Uint32s(key string, value ...uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint32p returns an attribute for the pointed value.
func Uint32p(key string, value *uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uint64 returns an attribute.
func Uint64(key string, value uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(value),
	}
}

// Uint64s returns an array attribute. The slice is copied.
func Uint64s(key string, value ...uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint64sValue(value),
	}
}

// Uint64p returns an attribute. A nil pointer is a nil value.
func Uint64p(key string, value *uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint64pValue(value),
	}
}

// Int returns an attribute.
func Int(key string, value int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

// Ints returns an array attribute. The slice is copied.
func Ints(key string, value ...int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Intp returns an attribute for the pointed value.
func Intp(key string, value *int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int8 returns an attribute.
func Int8(key string, value int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

// Int8s returns an array attribute. The slice is copied.
func Int8s(key string, value ...int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int8p returns an attribute for the pointed value.
func Int8p(key string, value *int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int16 returns an attribute.
func Int16(key string, value int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

// Int16s returns an array attribute. The slice is copied.
func Int16s(key string, value ...int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int16p returns an attribute for the pointed value.
func Int16p(key string, value *int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int32 returns an attribute.
func Int32(key string, value int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

// Int32s returns an array attribute. The slice is copied.
func Int32s(key string, value ...int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int32p returns an attribute for the pointed value.
func Int32p(key string, value *int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Int64 returns an attribute.
func Int64(key string, value int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(value),
	}
}

// Int64s returns an array attribute. The slice is copied.
func Int64s(key string, value ...int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Int64sValue(value),
	}
}

// Int64p returns an attribute. A nil pointer is a nil value.
func Int64p(key string, value *int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Int64pValue(value),
	}
}

// Float32 returns an attribute.
func Float32(key string, value float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Float64Value(float64(value)),
	}
}

// Float32s returns an array attribute. The slice is copied.
func Float32s(key string, value ...float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Float32p returns an attribute for the pointed value.
func Float32p(key string, value *float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Float64 returns an attribute.
func Float64(key string, value float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Float64Value(value),
	}
}

// Float64s returns an array attribute. The slice is copied.
func Float64s(key string, value ...float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Float64sValue(value),
	}
}

// Float64p returns an attribute. A nil pointer is a nil value.
func Float64p(key string, value *float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Float64pValue(value),
	}
}

// Complex64 returns an attribute.
func Complex64(key string, value complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(complex128(value)),
	}
}

// Complex64s returns an array attribute. The slice is copied.
func Complex64s(key string, value ...complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex64sValue(value),
	}
}

// Complex64p returns an attribute for the pointed value.
func Complex64p(key string, value *complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Complex128 returns an attribute.
func Complex128(key string, value complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(value),
	}
}

// Complex128s returns an array attribute. The slice is copied.
func Complex128s(key string, value ...complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex128sValue(value),
	}
}

// Complex128p returns an attribute. A nil pointer is a nil value.
func Complex128p(key string, value *complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex128pValue(value),
	}
}

// Uintptr returns an attribute.
func Uintptr(key string, value uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

// Uintptrs returns an array attribute. The slice is copied.
func Uintptrs(key string, value ...uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Uintptrp returns an attribute for the pointed value.
func Uintptrp(key string, value *uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Bytes returns a copied byte slice. The text and JSON encoders quote the bytes.
func Bytes(key string, value []byte) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BytesValue(value),
	}
}

// Duration returns an attribute.
func Duration(key string, value time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.DurationValue(value),
	}
}

// Durations returns an array attribute. The slice is copied.
func Durations(key string, value ...time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: DurationsValue(value),
	}
}

// Durationp returns an attribute. A nil pointer is a nil value.
func Durationp(key string, value *time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: DurationpValue(value),
	}
}

// Time returns an attribute.
func Time(key string, value time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.TimeValue(value),
	}
}

// Times returns an array attribute. The slice is copied.
func Times(key string, value ...time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: TimesValue(value),
	}
}

// Timep returns an attribute. A nil pointer is a nil value.
func Timep(key string, value *time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: TimepValue(value),
	}
}

// FormatTime returns an attribute formatted when it is logged.
func FormatTime(key, format string, value time.Time) slog.Attr {
	return slog.Attr{
		Key: key,
		Value: ClosureValue(func() any {
			return value.Format(format)
		}),
	}
}

// FormatTimes returns an attribute formatted when it is logged.
func FormatTimes(key, format string, value ...time.Time) slog.Attr {
	return slog.Attr{
		Key: key,
		Value: ClosureValue(func() any {
			times := make([]any, len(value))
			for idx, val := range value {
				times[idx] = val.Format(format)
			}

			return times
		}),
	}
}

// FormatTimep returns an attribute formatted when it is logged.
func FormatTimep(key, format string, value *time.Time) slog.Attr {
	isNill := value == nil

	return slog.Attr{
		Key: key,
		Value: ClosureValue(func() any {
			if isNill {
				return slog.AnyValue(nil)
			}

			return value.Format(format)
		}),
	}
}

// Error returns an error attribute. A nil error is an empty string.
func Error(key string, value error) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ErrorValue(value),
	}
}

// Errors returns an array of errors. The slice is copied.
func Errors(key string, value ...error) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ErrorsValue(value),
	}
}

// Groups returns a group attribute. The slice is copied.
func Groups(key string, value ...slog.Attr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: GroupValue(value...),
	}
}

// Valuer returns an attribute resolved through slog.LogValuer.
func Valuer(key string, value slog.LogValuer) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

// ValuerFn returns an attribute computed when it is logged.
func ValuerFn(key string, value ClosureFn) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ClosureValue(value),
	}
}
