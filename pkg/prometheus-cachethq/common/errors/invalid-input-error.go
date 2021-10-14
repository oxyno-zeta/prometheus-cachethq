package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewInvalidInputError(msg string) Error {
	return NewInvalidInputErrorWithExtensions(msg, nil)
}

func NewInvalidInputErrorWithError(err error) Error {
	return NewInvalidInputErrorWithExtensionsAndError(err, nil)
}

func NewInvalidInputErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewInvalidInputErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewInvalidInputErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "INVALID_INPUT"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusBadRequest,
	}
}
