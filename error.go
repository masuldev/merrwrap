package merrwrap

import "errors"

type WrapError struct {
	super  error
	origin error
}

func (e *WrapError) Super() error {
	return e.super
}

func (e *WrapError) Origin() error {
	return e.origin
}

func (e *WrapError) Wrap(err error) *WrapError {
	return &WrapError{super: err, origin: e}
}

func (e *WrapError) Error() string {
	if e.super == nil {
		return ""
	}

	msg := e.super.Error()
	if e.origin != nil {
		msg += " " + e.Origin().Error()
	}

	return msg
}

func Error(err error) *WrapError {
	if err == nil {
		return &WrapError{}
	}

	switch err.(type) {
	case *WrapError:
		return err.(*WrapError)
	default:
		return &WrapError{super: err}
	}
}

func (e *WrapError) Unwrap() error {
	return e.origin
}

func (e *WrapError) Is(target error) bool {
	return errors.Is(e.super, target)
}

func (e *WrapError) As(target interface{}) bool {
	return errors.As(e.super, target)
}
