package controller

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"

	"Mars/server/schemas"
	"Mars/shared/configure"
)

func Auth(c *atreugo.RequestCtx) error {
	token := c.Request.Header.Cookie("mars_token")
	if token == nil || unsafeConvert.StringPointer(token) != configure.Get().AuthToken {
		return c.JSONResponse(schemas.NewError("Authentication server is unavailable"), 403)
	}
	return c.Next()
}
