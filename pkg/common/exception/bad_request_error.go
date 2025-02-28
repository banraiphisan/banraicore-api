package exception

type BadRequestError struct {
	Message string
}

func (err BadRequestError) Error() string {
	return err.Message
}
