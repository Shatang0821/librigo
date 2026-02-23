package handler

import (
	"encoding/json"
	"errors"
	"librigo/internal/domain/apperror"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
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

	if status >= 500 {
		log.Printf("[ERROR] %d %s: %+v", status, code, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if encErr := json.NewEncoder(w).Encode(ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}); encErr != nil {
		// すでに Header や Status は書き込んでしまっているため、
		// ここではログを出力する以上のことはできませんが、異常を検知するために必須です
		log.Printf("[CRITICAL] Failed to encode error response: %v", encErr)
	}
}
