package log

import (
	"errors"
	"fmt"
)

type ErrCode int

// error code
const (
	ErrCodeHandler ErrCode = iota + 1
	ErrCodeService
	ErrCodeRepo
	ErrCodeMiddleware
	ErrCodeNotFound
)

var (
	ErrHandler    = customError{code: ErrCodeHandler, message: "handler error"}
	ErrService    = customError{code: ErrCodeService, message: "service error"}
	ErrRepo       = customError{code: ErrCodeRepo, message: "repository error"}
	ErrMiddleware = customError{code: ErrCodeMiddleware, message: "middleware error"}
	ErrNotFound   = customError{code: ErrCodeNotFound, message: "resource not found"}
)

type customError struct {
	code     ErrCode
	message  string
	original error
}

func (e customError) Error() string {
	if e.original != nil {
		return fmt.Sprintf("code=%d, error=%s", e.code, e.original.Error())
	}
	return fmt.Sprintf("code=%d, error=%s", e.code, e.message)
}

func (e customError) Unwrap() error {
	return e.original
}

func (e customError) Is(target error) bool {
	var customErr customError
	ok := errors.As(target, &customErr)
	if !ok {
		return ok
	}
	return e.code == customErr.code
}

func GetErrCode(err error) int {
	var customErr customError
	if errors.As(err, &customErr) {
		return int(customErr.code)
	}
	return -1 //default error code
}

func CustomError(custom customError, err error) error {
	var customErr customError
	if errors.As(err, &customErr) {
		return err
	}
	return customError{code: custom.code, message: custom.message, original: err}
}
