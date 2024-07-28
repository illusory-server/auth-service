package eerr

import "github.com/pkg/errors"

type (
	ErrorCode int

	Error struct {
		Code ErrorCode
		err  error
	}
)

func (e *Error) Error() string {
	return e.err.Error()
}

func New(code ErrorCode, msg string) error {
	return &Error{
		Code: code,
		err:  errors.New(msg),
	}
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	eErr, ok := err.(*Error)
	if ok {
		eErr.err = errors.Wrap(eErr.err, msg)
		return err
	}
	return &Error{
		Code: ErrInternal,
		err:  errors.Wrap(err, msg),
	}
}

func Code(err error) ErrorCode {
	if err == nil {
		return ErrInternal
	}
	if e, ok := err.(*Error); ok {
		return e.Code
	}
	return ErrInternal
}
