package gdel

import (
	"Mars/server/helper"
	"github.com/savsgio/atreugo/v11"
)

func Family(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()
	return nil
}
