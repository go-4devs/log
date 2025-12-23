package entry

import (
	"path/filepath"
	"runtime"
	"strconv"
)

func Caller(depth int, full bool) string {
	const offset = 3

	_, file, line, has := runtime.Caller(depth + offset)

	if !has {
		file, line = "???", 0
	}

	if !full && has {
		file = filepath.Base(file)
	}

	return file + ":" + strconv.Itoa(line)
}
