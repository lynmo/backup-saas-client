package swagger

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	error
	StatusCode() int
	Code() string
	Message() string
	OrigError() string
}

type ysError struct {
	YsapiError
}

// Unwrap to allow errors.Is and errors.As
func (e GenericSwaggerError) Unwrap() error {
	if e.body != nil {
		var ye = YsapiError{}
		err := json.Unmarshal(e.body, &ye)
		if err == nil {
			return ysError{YsapiError: ye}
		}
	}
	return nil
}

// satisfy error interface
func (e ysError) Error() string {
	if e.OrigError() != "" {
		return fmt.Sprintf("%s: %s\ncaused by: %s", e.Code(), e.Message(), e.OrigError())
	} else {
		return fmt.Sprintf("%s: %s", e.Code(), e.Message())
	}
}

// satisfy Error interface

func (e ysError) Code() string {
	return e.YsapiError.Code
}

func (e ysError) Message() string {
	return e.YsapiError.Message
}

func (e ysError) StatusCode() int {
	return int(e.YsapiError.StatusCode)
}

func (e ysError) OrigError() string {
	return e.YsapiError.OrigError
}
