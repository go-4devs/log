// Package buffer pools byte buffers used while formatting a log line.
package buffer

import "sync"

const bufferSize = 1024

// Buffer is a pooled byte buffer.
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

// New returns a buffer from the pool.
func New() *Buffer {
	buf, _ := bufPool.Get().(*Buffer)

	return buf
}

// Free returns the buffer to the pool when it is small enough.
func (b *Buffer) Free() {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}

// WriteString appends s.
func (b *Buffer) WriteString(s string) (int, error) {
	*b = append(*b, s...)

	return len(s), nil
}

// String returns the buffered text.
func (b *Buffer) String() string {
	return string(*b)
}
