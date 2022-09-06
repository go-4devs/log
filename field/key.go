package field

import (
	"time"
)

type Key string

//nolint:funlen,cyclop,gocyclo
func (k Key) Any(value interface{}) Field {
	switch val := value.(type) {
	case string:
		return k.String(val)
	case *string:
		return k.Stringp(val)
	case []string:
		return k.Strings(val...)
	case bool:
		return k.Bool(val)
	case *bool:
		return k.Boolp(val)
	case []bool:
		return k.Bools(val...)
	case int8:
		return k.Int8(val)
	case []int8:
		return k.Int8s(val...)
	case *int8:
		return k.Int8p(val)
	case int16:
		return k.Int16(val)
	case []int16:
		return k.Int16s(val...)
	case *int16:
		return k.Int16p(val)
	case int32:
		return k.Int32(val)
	case []int32:
		return k.Int32s(val...)
	case *int32:
		return k.Int32p(val)
	case int64:
		return k.Int64(val)
	case []int64:
		return k.Int64s(val...)
	case *int64:
		return k.Int64p(val)
	case uint:
		return k.Uint(val)
	case []uint:
		return k.Uints(val...)
	case *uint:
		return k.Uintp(val)
	case uint8:
		return k.Uint8(val)
	case *uint8:
		return k.Uint8p(val)
	case uint16:
		return k.Uint16(val)
	case []uint16:
		return k.Uint16s(val...)
	case *uint16:
		return k.Uint16p(val)
	case uint32:
		return k.Uint32(val)
	case []uint32:
		return k.Uint32s(val...)
	case *uint32:
		return k.Uint32p(val)
	case uint64:
		return k.Uint64(val)
	case []uint64:
		return k.Uint64s(val...)
	case *uint64:
		return k.Uint64p(val)
	case float32:
		return k.Float32(val)
	case []float32:
		return k.Float32s(val...)
	case *float32:
		return k.Float32p(val)
	case float64:
		return k.Float64(val)
	case []float64:
		return k.Float64s(val...)
	case *float64:
		return k.Float64p(val)
	case complex64:
		return k.Complex64(val)
	case []complex64:
		return k.Complex64s(val...)
	case *complex64:
		return k.Complex64p(val)
	case uintptr:
		return k.Uintptr(val)
	case []uintptr:
		return k.Uintptrs(val...)
	case *uintptr:
		return k.Uintptrp(val)
	case []byte:
		return k.Bytes(val)
	case time.Duration:
		return k.Dureation(val)
	case []time.Duration:
		return k.Dureations(val)
	case *time.Duration:
		return k.Dureationp(val)
	case time.Time:
		return k.Time(val)
	case []time.Time:
		return k.Times(val...)
	case *time.Time:
		return k.Timep(val)
	case error:
		return k.Error(val)
	case []error:
		return k.Errors(val...)
	}

	return Field{
		key: k,
		value: Value{
			value:    value,
			vtype:    TypeAny,
			numeric:  0,
			stringly: "",
		},
	}
}

func (k Key) String(value string) Field {
	return Field{
		key:   k,
		value: stringValue(value),
	}
}

func (k Key) Strings(value ...string) Field {
	return Field{
		key:   k,
		value: stringsValue(value),
	}
}

func (k Key) Stringp(value *string) Field {
	return Field{
		key:   k,
		value: stringpValue(value),
	}
}

func (k Key) Bool(value bool) Field {
	return Field{
		key:   k,
		value: boolValue(value),
	}
}

func (k Key) Bools(value ...bool) Field {
	return Field{
		key:   k,
		value: boolsValue(value),
	}
}

func (k Key) Boolp(value *bool) Field {
	return Field{
		key:   k,
		value: boolpValue(value),
	}
}

func (k Key) Int(value int) Field {
	return Field{
		key:   k,
		value: intValue(value),
	}
}

func (k Key) Ints(value ...int) Field {
	return Field{
		key:   k,
		value: intsValue(value),
	}
}

func (k Key) Intp(value *int) Field {
	return Field{
		key:   k,
		value: intpValue(value),
	}
}

func (k Key) Int8(value int8) Field {
	return Field{
		key:   k,
		value: int8Value(value),
	}
}

func (k Key) Int8s(value ...int8) Field {
	return Field{
		key:   k,
		value: int8sValue(value),
	}
}

func (k Key) Int8p(value *int8) Field {
	return Field{
		key:   k,
		value: int8pValue(value),
	}
}

func (k Key) Int16(value int16) Field {
	return Field{
		key:   k,
		value: int16Value(value),
	}
}

func (k Key) Int16s(value ...int16) Field {
	return Field{
		key:   k,
		value: int16sValue(value),
	}
}

func (k Key) Int16p(value *int16) Field {
	return Field{
		key:   k,
		value: int16pValue(value),
	}
}

func (k Key) Int32(value int32) Field {
	return Field{
		key:   k,
		value: int32Value(value),
	}
}

func (k Key) Int32s(value ...int32) Field {
	return Field{
		key:   k,
		value: int32sValue(value),
	}
}

func (k Key) Int32p(value *int32) Field {
	return Field{
		key:   k,
		value: int32pValue(value),
	}
}

func (k Key) Int64(value int64) Field {
	return Field{
		key:   k,
		value: int64Value(value),
	}
}

func (k Key) Int64s(value ...int64) Field {
	return Field{
		key:   k,
		value: int64sValue(value),
	}
}

func (k Key) Int64p(value *int64) Field {
	return Field{
		key:   k,
		value: int64pValue(value),
	}
}

func (k Key) Uint(value uint) Field {
	return Field{
		key:   k,
		value: uintValue(value),
	}
}

func (k Key) Uints(value ...uint) Field {
	return Field{
		key:   k,
		value: uintsValue(value),
	}
}

func (k Key) Uintp(value *uint) Field {
	return Field{
		key:   k,
		value: uintpValue(value),
	}
}

func (k Key) Uint8(value uint8) Field {
	return Field{
		key:   k,
		value: uint8Value(value),
	}
}

func (k Key) Uint8s(value ...uint8) Field {
	return Field{
		key:   k,
		value: uint8sValue(value),
	}
}

func (k Key) Uint8p(value *uint8) Field {
	return Field{
		key:   k,
		value: uint8pValue(value),
	}
}

func (k Key) Uint16(value uint16) Field {
	return Field{
		key:   k,
		value: uint16Value(value),
	}
}

func (k Key) Uint16s(value ...uint16) Field {
	return Field{
		key:   k,
		value: uint16sValue(value),
	}
}

func (k Key) Uint16p(value *uint16) Field {
	return Field{
		key:   k,
		value: uint16pValue(value),
	}
}

func (k Key) Uint32(value uint32) Field {
	return Field{
		key:   k,
		value: uint32Value(value),
	}
}

func (k Key) Uint32s(value ...uint32) Field {
	return Field{
		key:   k,
		value: uint32sValue(value),
	}
}

func (k Key) Uint32p(value *uint32) Field {
	return Field{
		key:   k,
		value: uint32pValue(value),
	}
}

func (k Key) Uint64(value uint64) Field {
	return Field{
		key:   k,
		value: uint64Value(value),
	}
}

func (k Key) Uint64s(value ...uint64) Field {
	return Field{
		key:   k,
		value: uint64sValue(value),
	}
}

func (k Key) Uint64p(value *uint64) Field {
	return Field{
		key:   k,
		value: uint64pValue(value),
	}
}

func (k Key) Float32(value float32) Field {
	return Field{
		key:   k,
		value: float32Value(value),
	}
}

func (k Key) Float32s(value ...float32) Field {
	return Field{
		key:   k,
		value: float32sValue(value),
	}
}

func (k Key) Float32p(value *float32) Field {
	return Field{
		key:   k,
		value: float32pValue(value),
	}
}

func (k Key) Float64(value float64) Field {
	return Field{
		key:   k,
		value: float64Value(value),
	}
}

func (k Key) Float64s(value ...float64) Field {
	return Field{
		key:   k,
		value: float64sValue(value),
	}
}

func (k Key) Float64p(value *float64) Field {
	return Field{
		key:   k,
		value: float64pValue(value),
	}
}

func (k Key) Complex64(value complex64) Field {
	return Field{
		key:   k,
		value: complex64Value(value),
	}
}

func (k Key) Complex64s(value ...complex64) Field {
	return Field{
		key:   k,
		value: complex64sValue(value),
	}
}

func (k Key) Complex64p(value *complex64) Field {
	return Field{
		key:   k,
		value: complex64pValue(value),
	}
}

func (k Key) Complex128(value complex128) Field {
	return Field{
		key:   k,
		value: complex128Value(value),
	}
}

func (k Key) Complex128s(value []complex128) Field {
	return Field{
		key:   k,
		value: complex128sValue(value),
	}
}

func (k Key) Complex128p(value *complex128) Field {
	return Field{
		key:   k,
		value: complex128pValue(value),
	}
}

func (k Key) Uintptr(value uintptr) Field {
	return Field{
		key:   k,
		value: uintptrValue(value),
	}
}

func (k Key) Uintptrs(value ...uintptr) Field {
	return Field{
		key:   k,
		value: uintptrsValue(value),
	}
}

func (k Key) Uintptrp(value *uintptr) Field {
	return Field{
		key:   k,
		value: uintptrpValue(value),
	}
}

func (k Key) Bytes(value []byte) Field {
	return Field{
		key:   k,
		value: bytesValue(value),
	}
}

func (k Key) Dureation(value time.Duration) Field {
	return Field{
		key:   k,
		value: durationValue(value),
	}
}

func (k Key) Dureations(value []time.Duration) Field {
	return Field{
		key:   k,
		value: durationsValue(value),
	}
}

func (k Key) Dureationp(value *time.Duration) Field {
	return Field{
		key:   k,
		value: durationpValue(value),
	}
}

func (k Key) Time(value time.Time) Field {
	return Field{
		key:   k,
		value: timeValue(value),
	}
}

func (k Key) Times(value ...time.Time) Field {
	return Field{
		key:   k,
		value: timesValue(value),
	}
}

func (k Key) Timep(value *time.Time) Field {
	return Field{
		key:   k,
		value: timepValue(value),
	}
}

func (k Key) FormatTime(format string, value time.Time) Field {
	return Field{
		key:   k,
		value: formatTimeValue(format, value),
	}
}

func (k Key) FormatTimes(format string, value ...time.Time) Field {
	return Field{
		key:   k,
		value: formatTimesValue(format, value),
	}
}

func (k Key) FormatTimep(format string, value *time.Time) Field {
	return Field{
		key:   k,
		value: formatTimepValue(format, value),
	}
}

func (k Key) Error(value error) Field {
	return Field{
		key:   k,
		value: errorValue(value),
	}
}

func (k Key) Errors(value ...error) Field {
	return Field{
		key:   k,
		value: errorsValue(value),
	}
}
