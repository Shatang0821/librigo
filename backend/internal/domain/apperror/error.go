package apperror

import "errors"

type ErrorType string

const (
	TypeNotFound        ErrorType = "NOT_FOUND"
	TypeConflict        ErrorType = "CONFLICT"
	TypeInvalid         ErrorType = "INVALID"
	TypeUnauthorized    ErrorType = "UNAUTHORIZED"
	TypeUnauthenticated ErrorType = "UNAUTHENTICATED"
)

type AppError struct {
	Err     error
	Code    string
	ErrType ErrorType
}

func New(err error, code string, t ErrorType) *AppError {
	return &AppError{
		Err:     err,
		Code:    code,
		ErrType: t,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

var (
	ErrInvalidJSON = New(errors.New("invalid JSON"), "INVALID_JSON", TypeInvalid)
	ErrNilInput    = New(errors.New("input is nil"), "NIL_INPUT", TypeInvalid)
)
