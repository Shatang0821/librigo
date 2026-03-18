package response

import (
	"encoding/json"
	"librigo/internal/errors"
	"net/http"
)

type APIResponse struct {
	Data  interface{}      `json:"data"`
	Error *errors.AppError `json:"error"`
}

// 成功レスポンスを返す関数
func WriteSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := APIResponse{
		Data:  data,
		Error: nil,
	}
	json.NewEncoder(w).Encode(res)
}

func HandleError(w http.ResponseWriter, err error) {
	appErr, ok := err.(*errors.AppError)

	if !ok {
		appErr = errors.ErrInternal
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.HTTPStatus)

	res := APIResponse{
		Data:  nil,
		Error: appErr,
	}

	json.NewEncoder(w).Encode(res)
}
