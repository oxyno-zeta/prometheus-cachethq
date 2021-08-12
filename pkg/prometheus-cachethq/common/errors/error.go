package errors

import "github.com/pkg/errors"

type stackTracerError interface {
	StackTrace() errors.StackTrace
}

// Error is the interface all common errors must implement.
type Error interface {
	error
	stackTracerError
	Extensions() map[string]interface{}
	StatusCode() int
}
