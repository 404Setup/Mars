package helper

import (
	"fmt"

	"github.com/savsgio/atreugo/v11"

	"Mars/server/schemas"
)

func HandleError(c *atreugo.RequestCtx, message string, statusCode int) error {
	return c.JSONResponse(schemas.NewError(message), statusCode)
}

func HandleInternalError(c *atreugo.RequestCtx, recover any) error {
	return HandleError(c, fmt.Sprintf("Internal error: %v", recover), 500)
}
