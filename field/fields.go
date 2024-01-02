package field

type Fields []Field

func (f Fields) Fields(fn func(Field) bool) {
	for idx := range f {
		if !fn(f[idx]) {
			return
		}
	}
}

func (f Fields) Any() any {
	fields := make(map[string]any)
	for idx := range f {
		fields[f[idx].Key] = f[idx].Value.Any()
	}

	return fields
}

func (f Fields) Append(fields ...Field) Fields {
	f = append(f, fields...)

	return f
}

func (f Fields) Set(idx int, field Field) {
	f[idx] = field
}

func (f Fields) Len() int {
	return len(f)
}
