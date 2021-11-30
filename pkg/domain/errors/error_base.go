package errors

import "fmt"

type DomainError struct{
	*ErrorBase
}


type AlreadyExistsDomainError struct{
	*ErrorBase
}

type DoesNotExistsDomainError struct{
	*ErrorBase
}



type ErrorBase struct {
	Code       string
	Message    string
	error
}

func NewError(code string, err error) *ErrorBase{
	return &ErrorBase{
		Code:       code,
		Message:    err.Error(),
	}
}

func (e ErrorBase) String() string{
	return fmt.Sprintf("%s %s", e.Code, e.Message)
}