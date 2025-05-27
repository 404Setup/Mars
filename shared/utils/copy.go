package utils

import (
	"io"
	_ "unsafe"
)

//go:linkname CopyZeroAlloc github.com/valyala/fasthttp.copyZeroAlloc
func CopyZeroAlloc(w io.Writer, r io.Reader) (int64, error)
