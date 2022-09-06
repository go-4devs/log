package field

import "time"

//nolint:interfacebloat
type Encoder interface {
	// Built-in types.
	AddArray(key string, value Value)
	AddAny(key string, value Value)
	AddNil(key string)
	AddBool(key string, value bool)
	AddBinary(key string, value []byte)
	AddInt(key string, value int)
	AddInt8(key string, value int8)
	AddInt16(key string, value int16)
	AddInt32(key string, value int32)
	AddInt64(key string, value int64)
	AddUint(key string, value uint)
	AddUint8(key string, value uint8)
	AddUint16(key string, value uint16)
	AddUint32(key string, value uint32)
	AddUint64(key string, value uint64)
	AddUintptr(key string, value uintptr)
	AddTime(key string, value time.Time)
	AddDuration(key string, value time.Duration)
	AddFloat32(key string, value float32)
	AddFloat64(key string, value float64)
	AddComplex64(key string, value complex64)
	AddComplex128(key string, value complex128)
	AddString(key, value string)
	AddError(key string, value error)
}
