package field

import (
	"fmt"
)

func NewEncoderText(opts ...func(*BaseEncoder)) BaseEncoder {
	opts = append([]func(*BaseEncoder){
		WithGropuConfig(0, 0, ' '),
		WithNullValue("<nil>"),
		WithDefaultValue(func(dst []byte, _ Encoder, val Value) []byte {
			return fmt.Appendf(dst, "%+v", val.Any())
		}),
	}, opts...)

	return NewEncoder(opts...)
}
