package customerror

import (
	"core/pkg/errorcode"
	"fmt"
)

type CustomError struct {
	Err           error
	IsBusinessErr bool
	Code          errorcode.ErrorCode
	Message       string
}

func (e *CustomError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("%v", e.Err)
	}
	return fmt.Sprintf("%s - %v", e.Message, e.Err)
}

type Opts struct {
	Code    errorcode.ErrorCode
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
		Message:       defaultOpts.Message,
	}
}
