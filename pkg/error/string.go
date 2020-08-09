package error

type stringError string

func (err stringError) Error() string {
	return string(err)
}

// NewString creates a very simple error without any context with the given message.
func NewString(message string) error {
	return stringError(message)
}
