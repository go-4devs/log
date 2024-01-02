package field

type LogValuer interface {
	LogValue() any
}

type ClosureFn func() any

func (v ClosureFn) LogValue() any {
	return v()
}
