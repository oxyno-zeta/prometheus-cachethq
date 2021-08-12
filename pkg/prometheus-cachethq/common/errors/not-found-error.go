package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewNotFoundError(msg string) Error {
	return NewNotFoundErrorWithExtensions(msg, nil)
}

func NewNotFoundErrorWithError(err error) Error {
	return NewNotFoundErrorWithExtensionsAndError(err, nil)
}

func NewNotFoundErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewNotFoundErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewNotFoundErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "NOT_FOUND"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusNotFound,
	}
}
