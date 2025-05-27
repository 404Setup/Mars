package hash

import (
	"errors"
	"fmt"
	"io"
	"mime"

	"github.com/3JoB/unsafeConvert"
	"github.com/minio/sha256-simd"

	"Mars/lib/http"
	"Mars/shared/utils"
	"Mars/shared/utils/json"
)

func FileHash(r io.Reader) (string, error) {
	h := sha256.New()
	defer h.Reset()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetFilenameAndHash(addr string, isRedirect bool) (string, string, error) {
	defer utils.GC()

	req, resp := http.Acquire(addr, "GET")
	resp.StreamBody = true
	defer http.Release(req, resp)
	if err := http.Do(req, resp); err != nil {
		return "", "", err
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 301 || resp.StatusCode() == 302 {
			if isRedirect {
				return "", "", errors.New("too many redirects")
			}
			locate := unsafeConvert.StringPointer(resp.Header.Peek("Location"))
			return GetFilenameAndHash(locate, true)
		}
		if json.JSON.Valid(resp.Body()) {
			return "", "", errors.New(unsafeConvert.StringPointer(resp.Body()))
		}
		return "", "", fmt.Errorf("unknown status code %d", resp.StatusCode())
	}
	filename := peekHeaderFilename(resp.Header.Peek("Content-Disposition"))
	if filename == "" {
		return "", "", errors.New("content-disposition is empty")
	}
	hash, err := FileHash(resp.BodyStream())
	if err != nil {
		return "", "", err
	}
	return filename, hash, nil
}

func peekHeaderFilename(d []byte) string {
	if d == nil {
		return ""
	}
	_, params, err := mime.ParseMediaType(unsafeConvert.StringPointer(d))
	if err != nil {
		return ""
	}
	return params["filename"]
}
