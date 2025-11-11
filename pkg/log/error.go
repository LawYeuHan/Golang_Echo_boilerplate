package log

import (
	"errors"
	"fmt"
)

// error code
const (
	ErrCodeHandler = iota + 1
	ErrCodeService
	ErrCodeRepo
	ErrCodeMiddleware
	ErrCodeNotFound
	ErrCodeDefault
)

type customError struct {
	code     int
	original error
}

func (e customError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.original.Error())
}

func (e customError) Unwrap() error {
	return e.original
}

func GetErrCode(err error) int {
	var customErr customError
	if errors.As(err, &customErr) {
		return customErr.code
	}
	return -1 //default error code
}

func CustomError(code int, err error) error {
	return customError{code: code, original: err}
}
