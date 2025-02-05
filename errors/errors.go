package errors

import (
	"fmt"

	pkgErr "github.com/pkg/errors"
)

func WithStack(err error) error {
	return pkgErr.WithStack(err)
}

func Wrap(err error, msg string) error {
	if msg == "" {
		return pkgErr.WithStack(err)
	}
	return pkgErr.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return pkgErr.Wrapf(err, format, args...)
}

type Terror struct {
	code  int32
	msg   string
	cause error
}

func New(code int32, format string, args ...interface{}) *Terror {
	return &Terror{
		code: code,
		msg:  fmt.Sprintf(format, args...),
	}
}

func (err *Terror) Clone(message string) *Terror {
	return New(err.code, message)
}

func (err *Terror) Clonef(format string, args ...interface{}) *Terror {
	return New(err.code, format, args...)
}

func (err *Terror) Code() int32 {
	return err.code
}

func (err *Terror) Msg() string {
	return err.msg
}

func (err *Terror) WithCause(cause error) *Terror {
	return &Terror{
		code:  err.code,
		msg:   err.msg,
		cause: cause,
	}
}

func (err *Terror) WithStack() error {
	return pkgErr.WithStack(err)
}

func (err *Terror) String() string {
	if err.cause == nil {
		return fmt.Sprintf("code:%v, msg:%v", err.code, err.msg)
	}
	return fmt.Sprintf("code:%v, msg:%v, cause:%v", err.code, err.msg, err.cause)
}

func (err *Terror) Error() string {
	return err.String()
}

func (err *Terror) Unwrap() error {
	if err.cause != nil {
		return err.cause
	}
	return nil
}

func (err *Terror) Is(target error) bool {
	if target == nil {
		return err.Code() == 0
	}

	if t, ok := target.(*Terror); ok {
		return t.Code() == err.Code()
	}

	if target = Unwrap(target); target == nil {
		return false
	}
	return err.Is(target)
}

func (err *Terror) As(target interface{}) bool {
	if t, ok := target.(**Terror); ok {
		*t = err
		return true
	}
	return false
}
