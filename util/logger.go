package util

type logger struct{}

func Trace(chan interface{}) {}

func Info(chan interface{}) {}

func Error(chan error) {}

type TraceLogger struct{}

type InfoLogger struct{}

type ErrorHandler struct {
	errorChan chan error
}

func (e *ErrorHandler) Handle(err error) {
	e.errorChan <- err
}
