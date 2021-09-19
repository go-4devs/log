package entry

import (
	"path/filepath"
	"runtime"
	"strconv"
)

func Caller(depth int, full bool) string {
	const offset = 4
	_, file, line, ok := runtime.Caller(depth + offset)

	if !ok {
		file, line = "???", 0
	}

	if !full && ok {
		file = filepath.Base(file)
	}

	return file + ":" + strconv.Itoa(line)
}
