package helpers

type ServiceError struct {
	StatusCode int
	Message    string
}

func (se *ServiceError) Error() string {
	return se.Message
}
