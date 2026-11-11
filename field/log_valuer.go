package field

import "log/slog"

type ClosureFn func() any

func (v ClosureFn) LogValue() slog.Value {
	return AnyValue(v())
}
