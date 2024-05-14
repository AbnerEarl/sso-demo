package tools

import (
	"path"
	"runtime"
)

func AbPathByCaller() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		return path.Dir(filename)
	}
	return ""
}
