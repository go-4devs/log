package entry

import "sync"

// nolint: gochecknoglobals
var pool = sync.Pool{
	New: func() interface{} {
		return New()
	},
}

//nolint: forcetypeassert
func Get() *Entry {
	e := pool.Get().(*Entry)
	e.Reset()

	return e
}

func Put(e *Entry) {
	pool.Put(e)
}
