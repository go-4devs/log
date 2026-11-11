package entry

import "sync"

//nolint:gochecknoglobals
var pool = sync.Pool{
	New: func() any {
		return New()
	},
}

func Get() *Entry {
	entry, ok := pool.Get().(*Entry)
	if !ok {
		entry = New()
	}

	entry.Reset()

	return entry
}

func Put(e *Entry) {
	pool.Put(e)
}
