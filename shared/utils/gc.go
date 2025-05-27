package utils

import (
	"runtime"

	"Mars/shared/configure"
)

func GC() {
	if configure.Get().ActiveGC {
		runtime.GC()
		if configure.Get().AllRecycled {
			runtime.GC()
			runtime.GC()
		}
	}
}
