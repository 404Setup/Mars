package json_se

import (
	"io"
	"log/slog"
	"runtime"

	"Mars/shared/configure"
	"Mars/shared/utils"
)

func NewJson() JSONSE {
	switch configure.Get().Json.Runtime {
	case "std2":
		return &jStd2{}
	case "gojson":
		return &goJson{}
	case "sonic":
		if runtime.GOARCH != "amd64" {
			slog.Warn("The sonic runtime is only available under x86_64/amd64, and has fallen back to std2.")
			return &jStd2{}
		}
		if !utils.HasAvx256 {
			slog.Warn("The CPU does not support AVX2, sonic is not available, and has fallen back to std2.")
			return &jStd2{}
		}
		return &jSonic{}
	case "sonnet":
		return &jSonnet{}
	case "segmentio":
		return &jSegmentIO{}
	default:
		return &jStd2{}
	}
}

type JSONSE interface {
	Decode(r io.Reader, v any) error
	Encode(w io.Writer, value any) error
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	Unmarshal(data []byte, v any) error
	Valid(data []byte) bool
}
