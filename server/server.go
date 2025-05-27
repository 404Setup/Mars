package server

import (
	"log/slog"
	"os"

	"github.com/3JoB/ulib/fsutil"
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"

	"Mars/server/schemas"
	"Mars/shared/configure"
	"Mars/shared/utils/json"
)

func Server() {
	mode := ""
	switch configure.Get().Server.ListenMode {
	case "tcp", "tcp6", "tcp4":
		mode = configure.Get().Server.ListenMode
	default:
		slog.Error("No server listen mode")
		os.Exit(-1)
	}

	fasthttp.SetBodySizePoolLimit(
		configure.Get().Server.RequestBodySize,
		configure.Get().Server.ResponseBodySize,
	)

	addr := ":" + unsafeConvert.IntToString(configure.Get().Server.ListenPort)
	if configure.Get().Server.ListenAddress != "" {
		addr = configure.Get().Server.ListenAddress + addr
	}

	conf := atreugo.Config{
		JSONMarshalFunc:    json.JSON.Encode,
		Network:            mode,
		Addr:               addr,
		MaxRequestBodySize: configure.Get().Server.MaxRequestBodySize * 1024 * 1024,
		NotFoundView: func(c *atreugo.RequestCtx) error {
			return c.JSONResponse(schemas.NewError("path not found"), 404)
		},
		ErrorView: func(c *atreugo.RequestCtx, err error, i int) {
			c.Response.SetStatusCode(i)
			_ = c.JSONResponse(schemas.NewError(err.Error()))
		},
		Name: "Tranic Mars",
	}

	if (configure.Get().Server.TLSInKey != "" && fsutil.IsFile(configure.Get().Server.TLSInKey)) &&
		(configure.Get().Server.TLSInCert != "" && fsutil.IsFile(configure.Get().Server.TLSInCert)) {
		slog.Info("SSL is enabled")
		conf.TLSEnable = true
		conf.CertFile = configure.Get().Server.TLSInCert
		conf.CertKey = configure.Get().Server.TLSInKey
	}

	a := atreugo.New(conf)
	Router(a)

	if err := a.ListenAndServe(); err != nil {
		panic(err)
	}
}
