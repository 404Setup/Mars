package json_se

import (
	"io"

	gjson "github.com/goccy/go-json"
)

type goJson struct {
}

func (*goJson) Decode(r io.Reader, v any) error {
	return gjson.NewDecoder(r).Decode(v)
}

func (*goJson) Encode(w io.Writer, value any) error {
	return gjson.NewEncoder(w).Encode(value)
}

func (*goJson) Marshal(v any) ([]byte, error) {
	return gjson.Marshal(v)
}

func (*goJson) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return gjson.MarshalIndent(v, prefix, indent)
}

func (*goJson) Unmarshal(data []byte, v any) error {
	return gjson.Unmarshal(data, v)
}

func (*goJson) Valid(data []byte) bool {
	return gjson.Valid(data)
}
