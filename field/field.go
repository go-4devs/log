package field

import (
	"log/slog"
	"time"
)

func Empty() slog.Attr {
	var attr slog.Attr

	return attr
}

func Any(key string, value any) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func String(key, value string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.StringValue(value),
	}
}

func Stringp(key string, value *string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: StringpValue(value),
	}
}

func Strings(key string, value ...string) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: StringsValue(value),
	}
}

func Bool(key string, value bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.BoolValue(value),
	}
}

func Bools(key string, value ...bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BoolsValue(value),
	}
}

func Boolp(key string, value *bool) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BoolpValue(value),
	}
}

func Uint(key string, value uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

func Uints(key string, value ...uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uintp(key string, value *uint) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint8(key string, value uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

func Uint8s(key string, value ...uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint8sValue(value),
	}
}

func Uint8p(key string, value *uint8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint16(key string, value uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

func Uint16s(key string, value ...uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint16p(key string, value *uint16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint32(key string, value uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

func Uint32s(key string, value ...uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint32p(key string, value *uint32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uint64(key string, value uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(value),
	}
}

func Uint64s(key string, value ...uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint64sValue(value),
	}
}

func Uint64p(key string, value *uint64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Uint64pValue(value),
	}
}

func Int(key string, value int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

func Ints(key string, value ...int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Intp(key string, value *int) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int8(key string, value int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

func Int8s(key string, value ...int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int8p(key string, value *int8) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int16(key string, value int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

func Int16s(key string, value ...int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int16p(key string, value *int16) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int32(key string, value int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(int64(value)),
	}
}

func Int32s(key string, value ...int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int32p(key string, value *int32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Int64(key string, value int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Int64Value(value),
	}
}

func Int64s(key string, value ...int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Int64sValue(value),
	}
}

func Int64p(key string, value *int64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Int64pValue(value),
	}
}

func Float32(key string, value float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Float64Value(float64(value)),
	}
}

func Float32s(key string, value ...float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Float32p(key string, value *float32) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Float64(key string, value float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Float64Value(value),
	}
}

func Float64s(key string, value ...float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Float64sValue(value),
	}
}

func Float64p(key string, value *float64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Float64pValue(value),
	}
}

func Complex64(key string, value complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(complex128(value)),
	}
}

func Complex64s(key string, value ...complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex64sValue(value),
	}
}

func Complex64p(key string, value *complex64) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Complex128(key string, value complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.AnyValue(value),
	}
}

func Complex128s(key string, value ...complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex128sValue(value),
	}
}

func Complex128p(key string, value *complex128) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: Complex128pValue(value),
	}
}

func Uintptr(key string, value uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(value)),
	}
}

func Uintptrs(key string, value ...uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Uintptrp(key string, value *uintptr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func Bytes(key string, value []byte) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: BytesValue(value),
	}
}

func Duration(key string, value time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.DurationValue(value),
	}
}

func Durations(key string, value ...time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: DurationsValue(value),
	}
}

func Durationp(key string, value *time.Duration) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: DurationpValue(value),
	}
}

func Time(key string, value time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: slog.TimeValue(value),
	}
}

func Times(key string, value ...time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: TimesValue(value),
	}
}

func Timep(key string, value *time.Time) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: TimepValue(value),
	}
}

func FormatTime(key, format string, value time.Time) slog.Attr {
	return slog.Attr{
		Key: key,
		Value: ClosureValue(func() any {
			return value.Format(format)
		}),
	}
}

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

func Error(key string, value error) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ErrorValue(value),
	}
}

func Errors(key string, value ...error) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ErrorsValue(value),
	}
}

func Groups(key string, value ...slog.Attr) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: GroupValue(value...),
	}
}

func Valuer(key string, value slog.LogValuer) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: AnyValue(value),
	}
}

func ValuerFn(key string, value ClosureFn) slog.Attr {
	return slog.Attr{
		Key:   key,
		Value: ClosureValue(value),
	}
}
