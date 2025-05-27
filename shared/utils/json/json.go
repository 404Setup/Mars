package json

import (
	"Mars/shared/utils/json/json_se"
)

var JSON json_se.JSONSE

func New() {
	JSON = json_se.NewJson()
}
