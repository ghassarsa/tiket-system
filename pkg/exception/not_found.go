package exception

type NotFound struct {
	Message string
}

func NewNotFoundError(message string) error {
	return &NotFound{Message: message}
}

func (e *NotFound) Error() string {
	return e.Message
}
