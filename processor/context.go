package processor

import (
	"context"
	"fmt"
)

func Context(key fmt.Stringer) Processor {
	return func(ctx context.Context) Field {
		return Field{key.String(), ctx.Value(key)}
	}
}
