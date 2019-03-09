package sample_error

import "fmt"

type SampleErrorCode int

const (
	UnknownError = iota
	UnAuthorized
	NotFoundError
	WrongPassword
	AlreadyExistError
	InternalError
)

type SampleError struct {
	Code SampleErrorCode
	Msg  string
}

func NewSampleError(code SampleErrorCode, msg string) *SampleError {
	return &SampleError{
		Code: code,
		Msg:  msg,
	}
}

func (err *SampleError) Error() string {
	switch err.Code {
	case UnAuthorized:
		return fmt.Sprintf("ERROR: UnAuthorizedError %s", err.Msg)
	case NotFoundError:
		return fmt.Sprintf("ERROR: NotFoundError %s", err.Msg)
	case WrongPassword:
		return fmt.Sprintf("ERROR: WrongPassword %s", err.Msg)
	case AlreadyExistError:
		return fmt.Sprintf("ERROR: AlreadyExistError %s", err.Msg)
	case InternalError:
		return fmt.Sprintf("ERROR: InternalError %s", err.Msg)
	}
	return fmt.Sprintf("ERROR: UnknownError")
}
