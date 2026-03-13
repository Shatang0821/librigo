package apperror

import "fmt"

type ErrorType string

const (
	TypeNotFound        ErrorType = "NOT_FOUND"
	TypeConflict        ErrorType = "CONFLICT"
	TypeInvalid         ErrorType = "INVALID"
	TypeUnauthorized    ErrorType = "UNAUTHORIZED"
	TypeUnauthenticated ErrorType = "UNAUTHENTICATED"
)

type ErrorDef struct {
	Code    string
	ErrType ErrorType
}

func New(code string, t ErrorType) *ErrorDef {
	return &ErrorDef{
		Code:    code,
		ErrType: t,
	}
}

func (d *ErrorDef) Error() string {
	return d.Code
}

// エラー定義をラップしてAppErrorを生成する
func (d *ErrorDef) Wrap(err error) *AppError {
	return &AppError{
		Def:   d,
		Cause: err,
	}
}

type AppError struct {
	Def   *ErrorDef
	Cause error
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Def.Code, e.Cause)
	}
	return e.Def.Code
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) Is(target error) bool {
	switch t := target.(type) {

	case *ErrorDef:
		return e.Def == t

	case *AppError:
		return e.Def == t.Def
	}

	return false
}

func (e *AppError) Code() string {
	return e.Def.Code
}

func (e *AppError) Type() ErrorType {
	return e.Def.ErrType
}

var (
	ErrInvalidJSON = New("INVALID_JSON", TypeInvalid)
	ErrNilInput    = New("NIL_INPUT", TypeInvalid)
)
