package exception

type BadRequest struct {
	Message string
}

func NewBadRequestError(message string) error {
	return &BadRequest{Message: message}
}

func (e *BadRequest) Error() string {
	return e.Message
}
