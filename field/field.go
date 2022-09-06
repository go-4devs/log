package field

import (
	"fmt"
	"time"
)

func Any(key string, value interface{}) Field {
	return Key(key).Any(value)
}

func String(key, value string) Field {
	return Key(key).String(value)
}

func Stringp(key string, value *string) Field {
	return Key(key).Stringp(value)
}

func Strings(key string, value ...string) Field {
	return Key(key).Strings(value...)
}

func Bool(key string, value bool) Field {
	return Key(key).Bool(value)
}

func Bools(key string, value ...bool) Field {
	return Key(key).Bools(value...)
}

func Boolp(key string, value *bool) Field {
	return Key(key).Boolp(value)
}

func Uint(key string, value uint) Field {
	return Key(key).Uint(value)
}

func Uints(key string, value ...uint) Field {
	return Key(key).Uints(value...)
}

func Uintp(key string, value *uint) Field {
	return Key(key).Uintp(value)
}

func Uint8(key string, value uint8) Field {
	return Key(key).Uint8(value)
}

func Uint8s(key string, value ...uint8) Field {
	return Key(key).Uint8s(value...)
}

func Uint8p(key string, value *uint8) Field {
	return Key(key).Uint8p(value)
}

func Uint16(key string, value uint16) Field {
	return Key(key).Uint16(value)
}

func Uint16s(key string, value ...uint16) Field {
	return Key(key).Uint16s(value...)
}

func Uint16p(key string, value *uint16) Field {
	return Key(key).Uint16p(value)
}

func Uint32(key string, value uint32) Field {
	return Key(key).Uint32(value)
}

func Uint32s(key string, value ...uint32) Field {
	return Key(key).Uint32s(value...)
}

func Uint32p(key string, value *uint32) Field {
	return Key(key).Uint32p(value)
}

func Uint64(key string, value uint64) Field {
	return Key(key).Uint64(value)
}

func Uint64s(key string, value ...uint64) Field {
	return Key(key).Uint64s(value...)
}

func Uint64p(key string, value *uint64) Field {
	return Key(key).Uint64p(value)
}

func Int(key string, value int) Field {
	return Key(key).Int(value)
}

func Ints(key string, value ...int) Field {
	return Key(key).Ints(value...)
}

func Intp(key string, value *int) Field {
	return Key(key).Intp(value)
}

func Int8(key string, value int8) Field {
	return Key(key).Int8(value)
}

func Int8s(key string, value ...int8) Field {
	return Key(key).Int8s(value...)
}

func Int8p(key string, value *int8) Field {
	return Key(key).Int8p(value)
}

func Int16(key string, value int16) Field {
	return Key(key).Int16(value)
}

func Int16s(key string, value ...int16) Field {
	return Key(key).Int16s(value...)
}

func Int16p(key string, value *int16) Field {
	return Key(key).Int16p(value)
}

func Int32(key string, value int32) Field {
	return Key(key).Int32(value)
}

func Int32s(key string, value ...int32) Field {
	return Key(key).Int32s(value...)
}

func Int32p(key string, value *int32) Field {
	return Key(key).Int32p(value)
}

func Int64(key string, value int64) Field {
	return Key(key).Int64(value)
}

func Int64s(key string, value ...int64) Field {
	return Key(key).Int64s(value...)
}

func Int64p(key string, value *int64) Field {
	return Key(key).Int64p(value)
}

func Float32(key string, value float32) Field {
	return Key(key).Float32(value)
}

func Float32s(key string, value ...float32) Field {
	return Key(key).Float32s(value...)
}

func Float32p(key string, value *float32) Field {
	return Key(key).Float32p(value)
}

func Float64(key string, value float64) Field {
	return Key(key).Float64(value)
}

func Float64s(key string, value ...float64) Field {
	return Key(key).Float64s(value...)
}

func Float64p(key string, value *float64) Field {
	return Key(key).Float64p(value)
}

func Complex64(key string, value complex64) Field {
	return Key(key).Complex64(value)
}

func Complex64s(key string, value ...complex64) Field {
	return Key(key).Complex64s(value...)
}

func Complex64p(key string, value *complex64) Field {
	return Key(key).Complex64p(value)
}

func Uintptr(key string, value uintptr) Field {
	return Key(key).Uintptr(value)
}

func Uintptrs(key string, value ...uintptr) Field {
	return Key(key).Uintptrs(value...)
}

func Uintptrp(key string, value *uintptr) Field {
	return Key(key).Uintptrp(value)
}

func Bytes(key string, value []byte) Field {
	return Key(key).Bytes(value)
}

func Duration(key string, value time.Duration) Field {
	return Key(key).Dureation(value)
}

func Durations(key string, value ...time.Duration) Field {
	return Key(key).Dureations(value)
}

func Durationp(key string, value *time.Duration) Field {
	return Key(key).Dureationp(value)
}

func Time(key string, value time.Time) Field {
	return Key(key).Time(value)
}

func Times(key string, value ...time.Time) Field {
	return Key(key).Times(value...)
}

func Timep(key string, value *time.Time) Field {
	return Key(key).Timep(value)
}

func FormatTime(key, format string, value time.Time) Field {
	return Key(key).FormatTime(format, value)
}

func FormatTimes(key, foramt string, value ...time.Time) Field {
	return Key(key).FormatTimes(foramt, value...)
}

func FormatTimep(key, foramt string, value *time.Time) Field {
	return Key(key).FormatTimep(foramt, value)
}

func Error(key string, value error) Field {
	return Key(key).Error(value)
}

func Errors(key string, value ...error) Field {
	return Key(key).Errors(value...)
}

// Field struct.
type Field struct {
	key   Key
	value Value
}

//nolint:gocyclo,cyclop
func (f Field) AddTo(enc Encoder) {
	key := string(f.key)

	switch {
	case f.value.IsArray():
		enc.AddArray(key, f.value)
	case f.value.IsNil():
		enc.AddNil(key)
	case f.value.IsBool():
		enc.AddBool(key, f.value.asBool())
	case f.value.IsBinary():
		enc.AddBinary(key, f.value.asBinary())
	case f.value.IsInt():
		enc.AddInt(key, f.value.asInt())
	case f.value.IsInt8():
		enc.AddInt8(key, f.value.asInt8())
	case f.value.IsInt16():
		enc.AddInt16(key, f.value.asInt16())
	case f.value.IsInt32():
		enc.AddInt32(key, f.value.asInt32())
	case f.value.IsInt64():
		enc.AddInt64(key, f.value.asInt64())
	case f.value.IsUint():
		enc.AddUint(key, f.value.asUint())
	case f.value.IsUint8():
		enc.AddUint8(key, f.value.asUint8())
	case f.value.IsUint16():
		enc.AddUint16(key, f.value.asUint16())
	case f.value.IsUint32():
		enc.AddUint32(key, f.value.asUint32())
	case f.value.IsUint64():
		enc.AddUint64(key, f.value.asUint64())
	case f.value.IsUintptr():
		enc.AddUintptr(key, f.value.asUintptr())
	case f.value.IsTime():
		enc.AddTime(key, f.value.asTime())
	case f.value.IsDuration():
		enc.AddDuration(key, f.value.asDuration())
	case f.value.IsFloat32():
		enc.AddFloat32(key, f.value.asFloat32())
	case f.value.IsFloat64():
		enc.AddFloat64(key, f.value.asFloat64())
	case f.value.IsComplex64():
		enc.AddComplex64(key, f.value.asComplex64())
	case f.value.IsComplex128():
		enc.AddComplex128(key, f.value.asComplex128())
	case f.value.IsString():
		enc.AddString(key, f.value.asString())
	case f.value.IsError():
		enc.AddError(key, f.value.asError())
	default:
		enc.AddAny(key, f.value)
	}
}

func (f Field) Type() Type {
	return f.value.vtype
}

func (f Field) Key() Key {
	return f.key
}

func (f Field) Value() Value {
	return f.value
}

func (f Field) AsInterface() interface{} {
	return f.value.AsInterface()
}

// String implent stringer.
func (f Field) String() string {
	return fmt.Sprintf("%s=%+v", f.key, f.value.AsInterface())
}
