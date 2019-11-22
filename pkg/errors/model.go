package errors

// BadInputErrorType Bad input error type
const BadInputErrorType = "BAD_INPUT_ERROR_TYPE"

// InternalServerErrorType Internal server error type
const InternalServerErrorType = "INTERNAL_SERVER_ERROR_TYPE"

// NotFoundErrorType Not found error type
const NotFoundErrorType = "NOT_FOUND_ERROR_TYPE"

// GeneralError General error
type GeneralError struct {
	ErrorType string
	Err       error
}

// Error for implementing error function
func (ge *GeneralError) Error() string {
	return ge.Err.Error()
}

// NewBadInputError New bad input error
func NewBadInputError(err error) *GeneralError {
	return &GeneralError{ErrorType: BadInputErrorType, Err: err}
}

// NewInternalServerError New internal server error
func NewInternalServerError(err error) *GeneralError {
	return &GeneralError{ErrorType: InternalServerErrorType, Err: err}
}

// NewNotFoundError New not found error
func NewNotFoundError(err error) *GeneralError {
	return &GeneralError{ErrorType: NotFoundErrorType, Err: err}
}
