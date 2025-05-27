package http

import (
	"github.com/valyala/fasthttp"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 TranicMars"

var Client = &fasthttp.Client{
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: false,
	Dial:                          dialer(),
}

func Acquire(uri, method string) (*fasthttp.Request, *fasthttp.Response) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	req.Header.Set("User-Agent", UserAgent)
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)
	return req, resp
}

func Release(req *fasthttp.Request, resp *fasthttp.Response) {
	_ = resp.CloseBodyStream()
	req.ReleaseBody(128)
	resp.ReleaseBody(1024)
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
}

func Do(req *fasthttp.Request, resp *fasthttp.Response) error {
	return Client.Do(req, resp)
}
