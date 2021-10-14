package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewConflictError(msg string) Error {
	return NewConflictErrorWithExtensions(msg, nil)
}

func NewConflictErrorWithError(err error) Error {
	return NewConflictErrorWithExtensionsAndError(err, nil)
}

func NewConflictErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewConflictErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewConflictErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "CONFLICT"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusConflict,
	}
}
