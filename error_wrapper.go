package error_wrapper

// ErrorWrapper use to wrap error, less error handle
type ErrorWrapper struct {
	err               error
	onErrorIsExecuted bool
}

// Do use to do something and it will return error or nil
func (w *ErrorWrapper) Do(f func() error) *ErrorWrapper {
	if w.err == nil {
		w.err = f()
	}
	return w
}

func (w *ErrorWrapper) GetErr() error {
	return w.err
}

func (w *ErrorWrapper) OnError(f func(err error)) *ErrorWrapper {
	if w.err != nil && w.onErrorIsExecuted != true {
		w.onErrorIsExecuted = true
		f(w.err)
	}
	return w
}
