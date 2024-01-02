package field

import (
	"encoding/json"
	"strconv"
)

func NewEncoderJSON(opts ...func(*BaseEncoder)) BaseEncoder {
	opts = append([]func(*BaseEncoder){
		WithAppendString(strconv.AppendQuote),
		WithDelimeter(':'),
		WithDefaultValue(func(dst []byte, e Encoder, val Value) []byte {
			js, err := json.Marshal(val.Any())
			if err != nil {
				return e.AppendValue(dst, ErrorValue(err))
			}

			return append(dst, js...)
		}),
	}, opts...)

	return NewEncoder(opts...)
}
