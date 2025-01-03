package ucerr

type UCError struct {
	message string
	ErrCode int
	debug   error
}

func New(message string, code int, err error) *UCError {
	return &UCError{
		message: message,
		ErrCode: code,
		debug:   err,
	}
}

func (e *UCError) Error() string {
	return e.message
}

func (e *UCError) Debug() error {
	return e.debug
}

func (e *UCError) Code() int {
	return e.ErrCode
}
