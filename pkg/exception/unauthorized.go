package exception

type Unauthorized struct {
	Message string
}

func NewUnauthorizedError(message string) error {
	return &Unauthorized{Message: message}
}

func (e *Unauthorized) Error() string
