package exception

type Internal struct {
	Message string
}

func NewInternalError(message string) error {
	return &Internal{Message: message}
}

func (e *Internal) Error() string
