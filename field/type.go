package field

type Type uint32

const (
	TypeAny        Type = 1 << iota // any
	TypeArray                       // array
	TypeNil                         // nil
	TypeString                      // string
	TypeBool                        // bool
	TypeInt                         // int
	TypeInt8                        // int8
	TypeInt16                       // int16
	TypeInt32                       // int32
	TypeInt64                       // int64
	TypeUint                        // uint
	TypeUint8                       // uint8
	TypeUint16                      // uint16
	TypeUint32                      // uint32
	TypeUint64                      // uint64
	TypeFloat32                     // float32
	TypeFloat64                     // float64
	TypeComplex64                   // complex64
	TypeComplex128                  // complex128
	TypeUintptr                     // uintptr
	TypeBinary                      // bytes
	TypeDuration                    // duration
	TypeTime                        // time
	TypeError                       // error
)

func (t Type) IsAny() bool {
	return t&TypeAny > 0
}

func (t Type) IsArray() bool {
	return t&TypeArray > 0
}

func (t Type) IsNil() bool {
	return t&TypeNil > 0
}

func (t Type) IsBool() bool {
	return t&TypeBool > 0
}

func (t Type) IsString() bool {
	return t&TypeString > 0
}

func (t Type) IsInt() bool {
	return t&TypeInt > 0
}

func (t Type) IsInt8() bool {
	return t&TypeInt8 > 0
}

func (t Type) IsInt16() bool {
	return t&TypeInt16 > 0
}

func (t Type) IsInt32() bool {
	return t&TypeInt32 > 0
}

func (t Type) IsInt64() bool {
	return t&TypeInt64 > 0
}

func (t Type) IsUint() bool {
	return t&TypeUint > 0
}

func (t Type) IsUint8() bool {
	return t&TypeUint8 > 0
}

func (t Type) IsUint16() bool {
	return t&TypeUint16 > 0
}

func (t Type) IsUint32() bool {
	return t&TypeUint32 > 0
}

func (t Type) IsUint64() bool {
	return t&TypeUint64 > 0
}

func (t Type) IsFloat32() bool {
	return t&TypeFloat32 > 0
}

func (t Type) IsFloat64() bool {
	return t&TypeFloat64 > 0
}

func (t Type) IsComplex64() bool {
	return t&TypeComplex64 > 0
}

func (t Type) IsComplex128() bool {
	return t&TypeComplex128 > 0
}

func (t Type) IsUintptr() bool {
	return t&TypeUintptr > 0
}

func (t Type) IsBinary() bool {
	return t&TypeBinary > 0
}

func (t Type) IsDuration() bool {
	return t&TypeDuration > 0
}

func (t Type) IsTime() bool {
	return t&TypeTime > 0
}

func (t Type) IsError() bool {
	return t&TypeError > 0
}
