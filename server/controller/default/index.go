package _default

import (
	"github.com/savsgio/atreugo/v11"

	schemas3 "Mars/shared/schemas"
)

func Index(c *atreugo.RequestCtx) error { return c.JSONResponse(schemas3.DefaultIndex, 200) }
