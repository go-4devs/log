package entry

import "sync"

//nolint:gochecknoglobals
var pool = sync.Pool{
	New: func() any {
		return New()
	},
}

func Get() *Entry {
	entry, _ := pool.Get().(*Entry)
	entry.Reset()

	return entry
}

func Put(e *Entry) {
	pool.Put(e)
}
