package processor

import (
	"runtime"

	"github.com/go-4devs/log"
)

// GoVersion add field by go version
func GoVersion(key string) log.Processor {
	return KeyValue(key, runtime.Version())
}
