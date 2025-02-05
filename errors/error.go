package errors

import (
	osErrors "errors"
	"fmt"
)

const (
	ServerErrorCode = 500
	ServerErrorMsg  = "internal error"

	UnknownErrorCode = 99999999
	UnknownErrorMsg  = "unknown error"
)

type Error interface {
	error
	Code() int32
	Msg() string
}

func Unwrap(err error) error {
	return osErrors.Unwrap(err)
}

func Is(err, target error) bool {
	return osErrors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return osErrors.As(err, target)
}

func Errorf(format string, a ...interface{}) error {
	return WithStack(fmt.Errorf(format, a...))
}

// IsBizError is the error about logic error, not rpc error, check by StatusCode != 0
func IsBizError(err error) bool {
	if err == nil {
		return false
	}

	if terr, ok := err.(*Terror); ok {
		if terr.cause != nil && terr.code != 0 {
			return IsBizError(terr.cause)
		}
		return terr.code != 0
	}

	if _, ok := err.(Error); ok {
		return true
	}

	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return false
	}
	return IsBizError(u.Unwrap())
}

func GetCode(err error) int32 {
	if err == nil {
		return 0
	}
	if v, ok := err.(Error); ok {
		return v.Code()
	}
	u, ok := err.(interface{ Unwrap() error })
	if !ok {
		return UnknownErrorCode
	}
	return GetCode(u.Unwrap())
}

func GetMsg(err error) string {
	if err == nil {
		return "success"
	}
	if v, ok := err.(Error); ok {
		return v.Msg()
	}
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return err.Error()
	}
	return GetMsg(u.Unwrap())
}
