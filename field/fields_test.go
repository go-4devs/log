package field_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitoa.ru/go-4devs/log/field"
)

func TestFields_Append(t *testing.T) {
	fields := field.Fields{field.Any("any", "value")}
	fields = fields.Append(field.String("string", "value"))

	require.Len(t, fields, 2)
}
