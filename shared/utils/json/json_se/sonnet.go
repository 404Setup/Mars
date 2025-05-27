package json_se

import (
	"io"

	"github.com/sugawarayuuta/sonnet"
)

type jSonnet struct {
}

func (*jSonnet) Decode(r io.Reader, v any) error {
	return sonnet.NewDecoder(r).Decode(v)
}

func (*jSonnet) Encode(w io.Writer, value any) error {
	return sonnet.NewEncoder(w).Encode(value)
}

func (*jSonnet) Marshal(v any) ([]byte, error) {
	return sonnet.Marshal(v)
}

func (*jSonnet) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return sonnet.MarshalIndent(v, prefix, indent)
}

func (*jSonnet) Unmarshal(data []byte, v any) error {
	return sonnet.Unmarshal(data, v)
}

func (*jSonnet) Valid(data []byte) bool {
	return sonnet.Valid(data)
}
