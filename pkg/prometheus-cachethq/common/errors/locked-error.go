package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewLockedError(msg string) Error {
	return NewLockedErrorWithExtensions(msg, nil)
}

func NewLockedErrorWithError(err error) Error {
	return NewLockedErrorWithExtensionsAndError(err, nil)
}

func NewLockedErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewLockedErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewLockedErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "LOCKED"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusLocked,
	}
}
