package merrwrap

type WrapError struct {
	current error
	child   error
}

func (e *WrapError) Current() error {
	return e.current
}

func (e *WrapError) Child() error {
	return e.child
}

func (e *WrapError) Wrap(err error) *WrapError {
	return &WrapError{current: err, child: e}
}

func (e *WrapError) Error() string {
	if e.current == nil {
		return ""
	}

	msg := e.current.Error()
	if e.child != nil {
		msg += " " + e.Child().Error()
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
		return &WrapError{current: err}
	}
}
