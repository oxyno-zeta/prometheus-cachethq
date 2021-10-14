package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewInternalServerError(msg string) Error {
	return NewInternalServerErrorWithExtensions(msg, nil)
}

func NewInternalServerErrorWithError(err error) Error {
	return NewInternalServerErrorWithExtensionsAndError(err, nil)
}

func NewInternalServerErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewInternalServerErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewInternalServerErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "INTERNAL_SERVER_ERROR"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusInternalServerError,
	}
}
