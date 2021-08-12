package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewUnauthorizedError(msg string) Error {
	return NewUnauthorizedErrorWithExtensions(msg, nil)
}

func NewUnauthorizedErrorWithError(err error) Error {
	return NewUnauthorizedErrorWithExtensionsAndError(err, nil)
}

func NewUnauthorizedErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewUnauthorizedErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewUnauthorizedErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "UNAUTHORIZED"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusUnauthorized,
	}
}
