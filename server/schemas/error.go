package schemas

type ErrorSchema struct {
	Error string `json:"error"`
}

func NewError(error string) *ErrorSchema {
	return &ErrorSchema{Error: error}
}

func NewErrorWithRaw(err error) *ErrorSchema {
	return &ErrorSchema{Error: err.Error()}
}

func (e *ErrorSchema) GetError() string {
	return e.Error
}
