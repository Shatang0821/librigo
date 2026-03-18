package errors

var (
	ErrInternal = &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    "internal server error",
		HTTPStatus: 500,
	}
)

// バリデーションエラーを生成関数
func NewValidationError(details []string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    "validation error",
		Details:    details,
		HTTPStatus: 400,
	}
}
