package apperror

type ErrorType string

const (
	TypeNotFound ErrorType = "NOT_FOUND"
	TypeConflict ErrorType = "CONFLICT"
	TypeInvalid  ErrorType = "INVALID"
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
