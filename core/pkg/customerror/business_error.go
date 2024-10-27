package customerror

import (
	"core/pkg/errorcode"
	"fmt"
)

type CustomError struct {
	Err           error
	IsBusinessErr bool
	Code          errorcode.ErrorCode
	Field         string
	Message       string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Business error on field '%s': %s - %s - %v", e.Field, e.Code.String(), e.Message, e.Err)
}

type Opts struct {
	Code    errorcode.ErrorCode
	Field   string
	Message string
}

func NewBusinessError(err error, opts ...Opts) error {
	defaultOpts := Opts{}

	if len(opts) > 0 {
		defaultOpts = opts[0]
	}

	return &CustomError{
		Err:           err,
		IsBusinessErr: true,
		Code:          defaultOpts.Code,
		Field:         defaultOpts.Field,
		Message:       defaultOpts.Message,
	}
}

func NewSystemError(err error, opts ...Opts) error {
	defaultOpts := Opts{}

	if len(opts) > 0 {
		defaultOpts = opts[0]
	}

	return &CustomError{
		Err:           err,
		IsBusinessErr: false,
		Code:          defaultOpts.Code,
		Field:         defaultOpts.Field,
		Message:       defaultOpts.Message,
	}
}
