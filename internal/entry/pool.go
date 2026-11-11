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

func Get() *logentry.Entry {
	ent, _ := entries.Get().(*logentry.Entry)
	ent.Reset()

	return ent
}

func Put(ent *logentry.Entry) {
	entries.Put(ent)
}
