package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewTooManyRequestsError(msg string) Error {
	return NewTooManyRequestsErrorWithExtensions(msg, nil)
}

func NewTooManyRequestsErrorWithError(err error) Error {
	return NewTooManyRequestsErrorWithExtensionsAndError(err, nil)
}

func NewTooManyRequestsErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewTooManyRequestsErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewTooManyRequestsErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "TOO_MANY_REQUESTS"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusTooManyRequests,
	}
}
