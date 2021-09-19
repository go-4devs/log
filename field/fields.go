package field

type Fields []Field

type MapField map[Key]Value

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

func (f Fields) AsMap() MapField {
	m := make(MapField, len(f))

	for _, field := range f {
		m[field.Key()] = field.Value()
	}

	return m
}
