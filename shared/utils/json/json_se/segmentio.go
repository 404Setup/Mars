package json_se

import (
	"io"

	"github.com/segmentio/encoding/json"
)

type jSegmentIO struct{}

func (*jSegmentIO) Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func (*jSegmentIO) Encode(w io.Writer, value any) error {
	return json.NewEncoder(w).Encode(value)
}

func (*jSegmentIO) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (*jSegmentIO) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (*jSegmentIO) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (*jSegmentIO) Valid(data []byte) bool {
	return json.Valid(data)
}
