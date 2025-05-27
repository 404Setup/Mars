//go:build amd64

package json_se

import (
	"io"

	"github.com/bytedance/sonic"

	"Mars/shared/configure"
)

type jSonic struct {
}

func (*jSonic) Decode(r io.Reader, v any) error {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.NewDecoder(r).Decode(v)
	}
	return sonic.ConfigDefault.NewDecoder(r).Decode(v)
}

func (*jSonic) Encode(w io.Writer, value any) error {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.NewEncoder(w).Encode(value)
	}
	return sonic.ConfigDefault.NewEncoder(w).Encode(value)
}

func (*jSonic) Marshal(v any) ([]byte, error) {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.Marshal(v)
	}
	return sonic.ConfigDefault.Marshal(v)
}

func (*jSonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.MarshalIndent(v, prefix, indent)
	}
	return sonic.ConfigDefault.MarshalIndent(v, prefix, indent)
}

func (*jSonic) Unmarshal(data []byte, v any) error {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.Unmarshal(data, v)
	}
	return sonic.ConfigDefault.Unmarshal(data, v)
}

func (*jSonic) Valid(data []byte) bool {
	if configure.Get().Json.Sonic.FastMode {
		return sonic.ConfigFastest.Valid(data)
	}
	return sonic.ConfigDefault.Valid(data)
}
