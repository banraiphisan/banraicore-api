package exception

type UnauthorizedError struct {
	Message string
}

func (err UnauthorizedError) Error() string {
	return err.Message
}
