package schemas

type ResultSchema struct {
	Result string `json:"result"`
}

func NewResult(result string) *ResultSchema {
	return &ResultSchema{Result: result}
}
