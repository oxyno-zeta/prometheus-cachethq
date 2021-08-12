package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewForbiddenError(msg string) Error {
	return NewForbiddenErrorWithExtensions(msg, nil)
}

func NewForbiddenErrorWithError(err error) Error {
	return NewForbiddenErrorWithExtensionsAndError(err, nil)
}

func NewForbiddenErrorWithExtensions(msg string, customExtensions map[string]interface{}) Error {
	return NewForbiddenErrorWithExtensionsAndError(errors.New(msg), customExtensions)
}

func NewForbiddenErrorWithExtensionsAndError(err error, customExtensions map[string]interface{}) Error {
	// Check if custom extensions exists
	if customExtensions == nil {
		customExtensions = map[string]interface{}{}
	}
	// Add code in custom extensions
	customExtensions["code"] = "FORBIDDEN"
	// Return new error
	return &GenericError{
		err:        errors.WithStack(err),
		ext:        customExtensions,
		statusCode: http.StatusForbidden,
	}
}
