package handler

import (
	"encoding/json"
	"errors"
	"librigo/internal/domain/apperror"
	"net/http"
)

type errorResponse struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RespondWithError はエラーを解析して適切なJSONレスポンスを返します
func RespondWithError(w http.ResponseWriter, err error) {
	var appErr *apperror.AppError
	var status int
	var code string
	var message string

	// 発生したエラーが AppError 型かどうかをチェック
	if errors.As(err, &appErr) {
		code = appErr.Code
		message = appErr.Error()

		// ErrorType を HTTP ステータスコードにマッピング
		switch appErr.ErrType {
		case apperror.TypeNotFound:
			status = http.StatusNotFound
		case apperror.TypeConflict:
			status = http.StatusConflict
		case apperror.TypeInvalid:
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
	} else {
		// AppError 以外（予期せぬエラー）の場合
		status = http.StatusInternalServerError
		code = "INTERNAL_SERVER_ERROR"
		message = "予期せぬエラーが発生しました"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{
		Error: errorDetail{
			Code:    code,
			Message: message,
		},
	})
}
