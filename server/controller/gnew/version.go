package gnew

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/utils/json"
)

func Version(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	var jar *schemas.NewVersionSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}

	ok, info := controller.CreateVersion(jar.Project, jar.Version, jar.Group)
	if ok {
		return c.JSONResponse(schemas.NewResult(info), 200)
	}

	return c.JSONResponse(schemas.NewError(info), 400)
}
