package errors

import "github.com/pkg/errors"

type GenericError struct {
	err        error
	ext        map[string]interface{}
	statusCode int
}

func (e *GenericError) Error() string {
	return e.err.Error()
}

func (e *GenericError) StackTrace() errors.StackTrace {
	// Cast internal error as a stack tracer error
	// nolint: errorlint // Ignore this because the aim is to catch stack trace error at first level
	if err2, ok := e.err.(stackTracerError); ok {
		return err2.StackTrace()
	}
	// Return nothing
	return nil
}

func (e *GenericError) Extensions() map[string]interface{} {
	return e.ext
}

func (e *GenericError) StatusCode() int {
	return e.statusCode
}
