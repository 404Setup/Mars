package gnew

import (
	"github.com/savsgio/atreugo/v11"

	"Mars/database/controller"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/utils/json"
)

func Build(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	var jar *schemas.NewBuildSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}

	if i, err := controller.CreateBuild(jar); err != nil {
		return c.JSONResponse(schemas.NewError(err.Error()), 400)
	} else {
		return c.JSONResponse(&schemas.NewBuildResult{
			Project: jar.Project,
			Version: jar.Version,
			Family:  jar.Family,
			BuildID: i,
		}, 200)
	}
}
