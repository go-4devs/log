package processor

import (
	"context"

	"github.com/go-4devs/log"
)

// KeyValue add field by const key value
func KeyValue(key string, value interface{}) log.Processor {
	return func(ctx context.Context) log.Field {
		return log.Field{Key: key, Value: value}
	}
}
