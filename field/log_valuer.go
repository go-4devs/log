package field

import "log/slog"

// ClosureFn is a function resolved when the attribute is logged.
type ClosureFn func() any

// LogValue resolves the closure.
func (v ClosureFn) LogValue() slog.Value {
	return AnyValue(v())
}
