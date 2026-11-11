// Package entry pools log entries for the root logger.
package entry

import (
	"sync"

	logentry "github.com/go-4devs/log/entry"
)

// shortcut: handler must not keep *Entry after return, copy the fields if it needs them later
//
//nolint:gochecknoglobals
var entries = sync.Pool{
	New: func() any {
		return logentry.New()
	},
}

// Get returns a reset entry from the pool.
func Get() *logentry.Entry {
	ent, _ := entries.Get().(*logentry.Entry)
	ent.Reset()

	return ent
}

// Put returns ent to the pool. Do not use ent after Put.
func Put(ent *logentry.Entry) {
	entries.Put(ent)
}
