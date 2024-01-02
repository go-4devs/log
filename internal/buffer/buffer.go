package buffer

import "sync"

const bufferSize = 1024

type Buffer []byte

// Having an initial size gives a dramatic speedup.
//
//nolint:gochecknoglobals
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, bufferSize)

		return (*Buffer)(&b)
	},
}

//nolint:forcetypeassert
func New() *Buffer {
	return bufPool.Get().(*Buffer)
}

func (b *Buffer) Free() {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}

func (b *Buffer) WriteString(s string) (int, error) {
	*b = append(*b, s...)

	return len(s), nil
}

func (b *Buffer) String() string {
	return string(*b)
}
