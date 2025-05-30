package hash

import (
	"errors"
	"fmt"
	"mime"

	fshash "github.com/3JoB/ulib/fsutil/hash"
	"github.com/3JoB/unsafeConvert"

	"Mars/lib/http"
	"Mars/shared/utils"
	"Mars/shared/utils/json"
)

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
	return filename, fshash.NewReader(resp.BodyStream(), &fshash.Opt{Crypt: fshash.SHA256}), nil
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
