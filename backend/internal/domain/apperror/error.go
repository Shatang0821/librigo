package apperror

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

func (d *ErrorDef) WithError(err error) *AppError {
	return &AppError{
		Err: err,
		Def: d,
	}
}

type AppError struct {
	Err error
	Def *ErrorDef
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Is(target error) bool {
	if t, ok := target.(*ErrorDef); ok {
		return e.Def.Code == t.Code // 自分のコードと定数のコードが一致するか？
	}

	return false
}

func (e *AppError) GetCode() string {
	return e.Def.Code
}

func (e *AppError) GetType() ErrorType {
	return e.Def.ErrType
}

var (
	ErrInvalidJSON = New("INVALID_JSON", TypeInvalid)
	ErrNilInput    = New("NIL_INPUT", TypeInvalid)
)
