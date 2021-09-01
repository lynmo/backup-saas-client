package swagger

import "fmt"

// Unwrap to allow errors.Is and errors.As
func (e GenericSwaggerError) Unwrap() error {
	if e.model != nil {
		if err, ok := e.model.(error); ok {
			return err
		}
	}
	return nil
}

// satisfy error interface
func (e YsapiError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
