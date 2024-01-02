package field

import (
	"encoding"
	"fmt"
)

func NewEncoderText(opts ...func(*BaseEncoder)) BaseEncoder {
	opts = append([]func(*BaseEncoder){
		WithGropuConfig(0, 0, ' '),
		WithNullValue("<nil>"),
		WithDefaultValue(func(dst []byte, enc Encoder, val Value) []byte {
			switch value := val.Any().(type) {
			case encoding.TextMarshaler:
				data, err := value.MarshalText()
				if err != nil {
					return enc.AppendValue(dst, ErrorValue(err))
				}

				return enc.AppendValue(dst, StringValue(string(data)))
			default:
				return fmt.Appendf(dst, "%+v", val.Any())
			}
		}),
	}, opts...)

	return NewEncoder(opts...)
}
