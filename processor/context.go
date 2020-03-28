package processor

import (
	"context"
	"fmt"

	"github.com/go-4devs/log"
)

// Context add field by context key
func Context(key fmt.Stringer) log.Processor {
	return func(ctx context.Context) log.Field {
		return log.Field{Key: key.String(), Value: ctx.Value(key)}
	}
}
