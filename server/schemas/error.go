package schemas

type ErrorSchema struct {
	Error string `json:"error"`
}

func NewError(error string) *ErrorSchema {
	return &ErrorSchema{Error: error}
}

func NewErrors(error error) *ErrorSchema {
	return &ErrorSchema{Error: error.Error()}
}

func (e *ErrorSchema) GetError() string {
	return e.Error
}
