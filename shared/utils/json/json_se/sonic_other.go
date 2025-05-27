//go:build !amd64

package json_se

import "io"

type jSonic struct {
}

func (*jSonic) Decode(r io.Reader, v any) error {
	return nil
}

func (*jSonic) Encode(w io.Writer, value any) error {
	return nil
}

func (*jSonic) Marshal(v any) ([]byte, error) {
	return nil, nil
}

func (*jSonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return nil, nil
}

func (*jSonic) Unmarshal(data []byte, v any) error {
	return nil
}

func (*jSonic) Valid(data []byte) bool { return false }
