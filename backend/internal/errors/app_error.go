package errors

type AppError struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Details    []string `json:"details,omitempty"`
	HTTPStatus int      `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}
