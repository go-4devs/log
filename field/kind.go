package field

import "fmt"

//go:generate stringer -type=Kind -linecomment -output=kind_string.go

type Kind int

const (
	KindAny        Kind = iota // any
	KindArray                  // array
	KindNil                    // nil
	KindString                 // string
	KindBool                   // bool
	KindInt64                  // int64
	KindUint64                 // uint64
	KindFloat32                // float32
	KindFloat64                // float64
	KindComplex128             // complex128
	KindBinary                 // bytes
	KindDuration               // duration
	KindTime                   // time
	KindError                  // error
	KindGroup                  // group
	KindClosure                // closure
)

func (l Kind) MarshalJSON() ([]byte, error) {
	return []byte("\"" + l.String() + "\""), nil
}

func (l *Kind) UnmarshalJSON(in []byte) error {
	return l.UnmarshalText(in[1 : len(in)-1])
}

func (l Kind) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

//nolint:gocyclo,cyclop,funlen
func (l *Kind) UnmarshalText(in []byte) error {
	switch string(in) {
	case KindAny.String():
		*l = KindAny

		return nil
	case KindArray.String():
		*l = KindArray

		return nil
	case KindNil.String():
		*l = KindNil

		return nil
	case KindString.String():
		*l = KindString

		return nil
	case KindBool.String():
		*l = KindBool

		return nil
	case KindInt64.String():
		*l = KindInt64

		return nil
	case KindUint64.String():
		*l = KindUint64

		return nil
	case KindFloat32.String():
		*l = KindFloat32

		return nil
	case KindFloat64.String():
		*l = KindFloat64

		return nil
	case KindComplex128.String():
		*l = KindComplex128

		return nil
	case KindBinary.String():
		*l = KindBinary

		return nil
	case KindDuration.String():
		*l = KindDuration

		return nil
	case KindTime.String():
		*l = KindTime

		return nil
	case KindError.String():
		*l = KindError

		return nil
	case KindGroup.String():
		*l = KindGroup

		return nil
	case KindClosure.String():
		*l = KindClosure

		return nil
	default:
		return fmt.Errorf("%w:filed(%v)", ErrUndefined, string(in))
	}
}

func (l Kind) MarshalBinary() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Kind) UnmarshalBinary(in []byte) error {
	return l.UnmarshalText(in)
}
