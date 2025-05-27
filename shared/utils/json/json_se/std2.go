package json_se

import (
	j1 "encoding/json"
	"io"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type jStd2 struct {
}

func (*jStd2) Decode(r io.Reader, v any) error {
	return json.UnmarshalDecode(jsontext.NewDecoder(r), v)
}

func (*jStd2) Encode(w io.Writer, value any) error {
	return json.MarshalEncode(jsontext.NewEncoder(w), value)
}

func (*jStd2) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (*jStd2) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.Marshal(v, jsontext.WithIndent(indent), jsontext.WithIndentPrefix(prefix))
}

func (*jStd2) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (*jStd2) Valid(data []byte) bool {
	return j1.Valid(data)
}
