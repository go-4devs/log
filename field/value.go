package field

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

type Value struct {
	vtype    Type
	numeric  uint64
	stringly string
	value    interface{}
}

func (v Value) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(v.AsInterface())
	if err != nil {
		return nil, fmt.Errorf("marshal err: %w", err)
	}

	return b, nil
}

//nolint: gocyclo,gomnd,cyclop
func (v Value) String() string {
	switch {
	case v.vtype.IsArray(), v.vtype.IsAny():
		return fmt.Sprintf("%+v", v.AsInterface())
	case v.vtype.IsNil():
		return "<nil>"
	case v.vtype.IsString():
		return v.asString()
	case v.vtype.IsBool():
		return strconv.FormatBool(v.asBool())
	case v.vtype.IsInt(), v.vtype.IsInt8(), v.vtype.IsInt16(), v.vtype.IsInt32():
		return strconv.Itoa(v.asInt())
	case v.vtype.IsInt64():
		return strconv.FormatInt(v.asInt64(), 10)
	case v.vtype.IsUint(), v.vtype.IsUint8(), v.vtype.IsUint16(), v.vtype.IsUint32(), v.vtype.IsUint64():
		return strconv.FormatUint(v.asUint64(), 10)
	case v.vtype.IsFloat64():
		return strconv.FormatFloat(v.asFloat64(), 'g', -1, 64)
	case v.vtype.IsFloat32():
		return strconv.FormatFloat(float64(v.asFloat32()), 'g', -1, 32)
	case v.vtype.IsComplex128():
		return strconv.FormatComplex(v.asComplex128(), 'g', -1, 128)
	case v.vtype.IsComplex64():
		return strconv.FormatComplex(complex128(v.asComplex64()), 'g', -1, 64)
	case v.vtype.IsBinary():
		return string(v.asBinary())
	case v.vtype.IsDuration():
		return v.asDuration().String()
	case v.vtype.IsTime():
		return v.asTime().Format(v.asString())
	case v.vtype.IsError():
		return v.asError().Error()
	}

	return fmt.Sprintf("%+v", v.AsInterface())
}

//nolint: gocyclo,cyclop
func (v Value) AsInterface() interface{} {
	switch {
	case v.vtype.IsArray():
		return v.value
	case v.vtype.IsNil():
		return nil
	case v.vtype.IsString():
		return v.asString()
	case v.vtype.IsBool():
		return v.asBool()
	case v.vtype.IsInt():
		return v.asInt()
	case v.vtype.IsInt8():
		return v.asInt8()
	case v.vtype.IsInt16():
		return v.asInt16()
	case v.vtype.IsInt32():
		return v.asInt32()
	case v.vtype.IsInt64():
		return v.asInt64()
	case v.vtype.IsUint():
		return v.asUint()
	case v.vtype.IsUint8():
		return v.asUint8()
	case v.vtype.IsUint16():
		return v.asUint16()
	case v.vtype.IsUint32():
		return v.asUint32()
	case v.vtype.IsUint64():
		return v.asUint64()
	case v.vtype.IsFloat32():
		return v.asFloat32()
	case v.vtype.IsFloat64():
		return v.asFloat64()
	case v.vtype.IsComplex64():
		return v.asComplex64()
	case v.vtype.IsComplex128():
		return v.asComplex128()
	case v.vtype.IsUintptr():
		return v.asUintptr()
	case v.vtype.IsBinary():
		return v.asBinary()
	case v.vtype.IsDuration():
		return v.asDuration()
	case v.vtype.IsTime():
		return v.asTime()
	case v.vtype.IsError():
		return v.asError()
	}

	return v.value
}

func (v Value) IsArray() bool {
	return v.vtype.IsArray()
}

func (v Value) IsNil() bool {
	return v.vtype.IsNil()
}

func (v Value) IsString() bool {
	return v.vtype.IsString()
}

func (v Value) IsBool() bool {
	return v.vtype.IsBool()
}

func (v Value) IsInt() bool {
	return v.vtype.IsInt()
}

func (v Value) IsInt8() bool {
	return v.vtype.IsInt8()
}

func (v Value) IsInt16() bool {
	return v.vtype.IsInt16()
}

func (v Value) IsInt32() bool {
	return v.vtype.IsInt32()
}

func (v Value) IsInt64() bool {
	return v.vtype.IsInt64()
}

func (v Value) IsUint() bool {
	return v.vtype.IsUint()
}

func (v Value) IsUint8() bool {
	return v.vtype.IsUint8()
}

func (v Value) IsUint16() bool {
	return v.vtype.IsUint16()
}

func (v Value) IsUint32() bool {
	return v.vtype.IsUint32()
}

func (v Value) IsUint64() bool {
	return v.vtype.IsUint64()
}

func (v Value) IsFloat32() bool {
	return v.vtype.IsFloat32()
}

func (v Value) IsFloat64() bool {
	return v.vtype.IsFloat64()
}

func (v Value) IsComplex64() bool {
	return v.vtype.IsComplex64()
}

func (v Value) IsComplex128() bool {
	return v.vtype.IsComplex128()
}

func (v Value) IsUintptr() bool {
	return v.vtype.IsUintptr()
}

func (v Value) IsBinary() bool {
	return v.vtype.IsBinary()
}

func (v Value) IsDuration() bool {
	return v.vtype.IsDuration()
}

func (v Value) IsTime() bool {
	return v.vtype.IsTime()
}

func (v Value) IsError() bool {
	return v.vtype.IsError()
}

func (v Value) asString() string {
	return v.stringly
}

func (v Value) asBool() bool {
	return v.numeric == 1
}

func (v Value) asInt() int {
	return int(v.numeric)
}

func (v Value) asInt8() int8 {
	return int8(v.numeric)
}

func (v Value) asInt16() int16 {
	return int16(v.numeric)
}

func (v Value) asInt32() int32 {
	return int32(v.numeric)
}

func (v Value) asInt64() int64 {
	return int64(v.numeric)
}

func (v Value) asUint() uint {
	return uint(v.numeric)
}

func (v Value) asUint8() uint8 {
	return uint8(v.numeric)
}

func (v Value) asUint16() uint16 {
	return uint16(v.numeric)
}

func (v Value) asUint32() uint32 {
	return uint32(v.numeric)
}

func (v Value) asUint64() uint64 {
	return v.numeric
}

func (v Value) asFloat32() float32 {
	return math.Float32frombits(uint32(v.numeric))
}

func (v Value) asFloat64() float64 {
	return math.Float64frombits(v.numeric)
}

func (v Value) asComplex64() complex64 {
	cmplex, _ := v.value.(complex64)

	return cmplex
}

func (v Value) asComplex128() complex128 {
	cmplex, _ := v.value.(complex128)

	return cmplex
}

func (v Value) asUintptr() uintptr {
	val, _ := v.value.(uintptr)

	return val
}

func (v Value) asBinary() []byte {
	bytes, _ := v.value.([]byte)

	return bytes
}

func (v Value) asDuration() time.Duration {
	duration, _ := v.value.(time.Duration)

	return duration
}

func (v Value) asTime() time.Time {
	value, _ := v.value.(time.Time)

	return value
}

func (v Value) asError() error {
	err, _ := v.value.(error)

	return err
}

func nilValue(t Type) Value {
	return Value{
		vtype:    t | TypeNil,
		value:    nil,
		numeric:  0,
		stringly: "",
	}
}

func stringValue(v string) Value {
	return Value{
		stringly: v,
		vtype:    TypeString,
		numeric:  0,
		value:    nil,
	}
}

func stringsValue(v []string) Value {
	return Value{
		value:    v,
		vtype:    TypeString | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func stringpValue(v *string) Value {
	if v != nil {
		return stringValue(*v)
	}

	return nilValue(TypeString)
}

func boolValue(b bool) Value {
	if b {
		return Value{
			numeric:  1,
			vtype:    TypeBool,
			value:    nil,
			stringly: "",
		}
	}

	return Value{
		vtype:    TypeBool,
		value:    nil,
		numeric:  0,
		stringly: "",
	}
}

func boolsValue(b []bool) Value {
	return Value{
		value:    b,
		vtype:    TypeBool | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func boolpValue(b *bool) Value {
	if b != nil {
		return boolValue(*b)
	}

	return nilValue(TypeBool)
}

func intValue(i int) Value {
	return Value{
		vtype:    TypeInt,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func intsValue(i []int) Value {
	return Value{
		value:    i,
		vtype:    TypeInt | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func intpValue(in *int) Value {
	if in != nil {
		return intValue(*in)
	}

	return nilValue(TypeInt)
}

func int8Value(i int8) Value {
	return Value{
		vtype:    TypeInt8,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func int8sValue(i []int8) Value {
	return Value{
		value:    i,
		vtype:    TypeInt8 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func int8pValue(in *int8) Value {
	if in != nil {
		return int8Value(*in)
	}

	return nilValue(TypeInt8)
}

func int16Value(i int16) Value {
	return Value{
		vtype:    TypeInt16,
		numeric:  uint64(i),
		value:    0,
		stringly: "",
	}
}

func int16sValue(i []int16) Value {
	return Value{
		value:    i,
		vtype:    TypeInt16 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func int16pValue(in *int16) Value {
	if in != nil {
		return int16Value(*in)
	}

	return nilValue(TypeInt16)
}

func int32Value(i int32) Value {
	return Value{
		vtype:    TypeInt32,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func int32sValue(i []int32) Value {
	return Value{
		value:    i,
		vtype:    TypeInt32 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func int32pValue(in *int32) Value {
	if in != nil {
		return int32Value(*in)
	}

	return nilValue(TypeInt32)
}

func int64Value(i int64) Value {
	return Value{
		vtype:    TypeInt64,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func int64sValue(i []int64) Value {
	return Value{
		value:    i,
		vtype:    TypeInt64 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func int64pValue(in *int64) Value {
	if in != nil {
		return int64Value(*in)
	}

	return nilValue(TypeInt64)
}

func uintValue(i uint) Value {
	return Value{
		vtype:    TypeUint,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func uintsValue(i []uint) Value {
	return Value{
		value:    i,
		vtype:    TypeUint | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func uintpValue(in *uint) Value {
	if in != nil {
		return uintValue(*in)
	}

	return nilValue(TypeUint)
}

func uint8Value(i uint8) Value {
	return Value{
		vtype:    TypeUint8,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func uint8sValue(i []uint8) Value {
	return Value{
		value:    i,
		vtype:    TypeUint8 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func uint8pValue(in *uint8) Value {
	if in != nil {
		return uint8Value(*in)
	}

	return nilValue(TypeUint8)
}

func uint16Value(i uint16) Value {
	return Value{
		vtype:    TypeUint16,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func uint16sValue(i []uint16) Value {
	return Value{
		value:    i,
		vtype:    TypeUint16 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func uint16pValue(in *uint16) Value {
	if in != nil {
		return uint16Value(*in)
	}

	return nilValue(TypeUint16)
}

func uint32Value(i uint32) Value {
	return Value{
		vtype:    TypeUint32,
		numeric:  uint64(i),
		value:    nil,
		stringly: "",
	}
}

func uint32sValue(i []uint32) Value {
	return Value{
		value:    i,
		vtype:    TypeUint32 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func uint32pValue(in *uint32) Value {
	if in != nil {
		return uint32Value(*in)
	}

	return nilValue(TypeUint32)
}

func uint64Value(i uint64) Value {
	return Value{
		vtype:    TypeUint64,
		numeric:  i,
		value:    nil,
		stringly: "",
	}
}

func uint64sValue(i []uint64) Value {
	return Value{
		value:    i,
		vtype:    TypeUint64 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func uint64pValue(in *uint64) Value {
	if in != nil {
		return uint64Value(*in)
	}

	return nilValue(TypeUint64)
}

func float32Value(i float32) Value {
	return Value{
		vtype:    TypeFloat32,
		numeric:  uint64(math.Float32bits(i)),
		value:    nil,
		stringly: "",
	}
}

func float32sValue(i []float32) Value {
	return Value{
		value:    i,
		vtype:    TypeFloat32 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func float32pValue(in *float32) Value {
	if in != nil {
		return float32Value(*in)
	}

	return nilValue(TypeFloat32)
}

func float64Value(i float64) Value {
	return Value{
		vtype:    TypeFloat64,
		numeric:  math.Float64bits(i),
		value:    nil,
		stringly: "",
	}
}

func float64sValue(i []float64) Value {
	return Value{
		value:    i,
		vtype:    TypeFloat64 | TypeArray,
		numeric:  0,
		stringly: "",
	}
}

func float64pValue(in *float64) Value {
	if in != nil {
		return float64Value(*in)
	}

	return nilValue(TypeFloat64)
}

func complex64Value(in complex64) Value {
	return Value{
		vtype:    TypeComplex64,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func complex64sValue(in []complex64) Value {
	return Value{
		vtype:    TypeComplex64 | TypeArray,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func complex64pValue(in *complex64) Value {
	if in != nil {
		return complex64Value(*in)
	}

	return nilValue(TypeComplex64)
}

func complex128Value(in complex128) Value {
	return Value{
		vtype:    TypeComplex64,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func complex128sValue(in []complex128) Value {
	return Value{
		vtype:    TypeComplex128 | TypeArray,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func complex128pValue(in *complex128) Value {
	if in != nil {
		return complex128Value(*in)
	}

	return nilValue(TypeComplex128)
}

func uintptrValue(in uintptr) Value {
	return Value{
		vtype:    TypeUintptr,
		numeric:  0,
		stringly: "",
		value:    in,
	}
}

func uintptrsValue(in []uintptr) Value {
	return Value{
		vtype:    TypeUintptr | TypeArray,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func uintptrpValue(in *uintptr) Value {
	if in != nil {
		return uintptrValue(*in)
	}

	return nilValue(TypeUintptr)
}

func bytesValue(in []byte) Value {
	return Value{
		vtype:    TypeBinary,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func durationValue(in time.Duration) Value {
	return Value{
		vtype:    TypeDuration,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func durationsValue(in []time.Duration) Value {
	return Value{
		vtype:    TypeDuration | TypeArray,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}

func durationpValue(in *time.Duration) Value {
	if in != nil {
		return durationValue(*in)
	}

	return nilValue(TypeDuration)
}

func timeValue(in time.Time) Value {
	return formatTimeValue(time.RFC3339, in)
}

func timesValue(in []time.Time) Value {
	return formatTimesValue(time.RFC3339, in)
}

func timepValue(in *time.Time) Value {
	return formatTimepValue(time.RFC3339, in)
}

func formatTimeValue(format string, in time.Time) Value {
	return Value{
		vtype:    TypeTime,
		value:    in,
		stringly: format,
		numeric:  0,
	}
}

func formatTimesValue(format string, in []time.Time) Value {
	return Value{
		vtype:    TypeTime | TypeArray,
		value:    in,
		stringly: format,
		numeric:  0,
	}
}

func formatTimepValue(format string, in *time.Time) Value {
	if in != nil {
		return formatTimeValue(format, *in)
	}

	return nilValue(TypeTime)
}

func errorValue(in error) Value {
	if in != nil {
		return Value{
			vtype:    TypeError,
			value:    in,
			numeric:  0,
			stringly: "",
		}
	}

	return nilValue(TypeError)
}

func errorsValue(in []error) Value {
	return Value{
		vtype:    TypeError | TypeArray,
		value:    in,
		numeric:  0,
		stringly: "",
	}
}
