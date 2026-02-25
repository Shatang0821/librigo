package handler_test

import (
	"encoding/json"
	"errors"
	"librigo/internal/domain/apperror"
	"librigo/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	tests := map[string]struct {
		err        error
		wantStatus int
		wantCode   string
	}{
		"AppError: NotFound": {
			err:        apperror.New(errors.New("not found"), "ITEM_NOT_FOUND", apperror.TypeNotFound),
			wantStatus: http.StatusNotFound,
			wantCode:   "ITEM_NOT_FOUND",
		},
		"AppError: Conflict": {
			err:        apperror.New(errors.New("already exists"), "DUPLICATE", apperror.TypeConflict),
			wantStatus: http.StatusConflict,
			wantCode:   "DUPLICATE",
		},
		"想定外のエラー": {
			err:        errors.New("db error"),
			wantStatus: http.StatusInternalServerError,
			wantCode:   "INTERNAL_SERVER_ERROR",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// 1. Recorderとリクエストを準備
			w := httptest.NewRecorder()

			// 2. テスト対象の実行
			handler.RespondWithError(w, tt.err)

			// 3. ステータスコードの検証
			if w.Code != tt.wantStatus {
				t.Errorf("status code: got %d, want %d", w.Code, tt.wantStatus)
			}

			// 4. JSONボディの検証
			var res handler.ErrorResponse
			if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if res.Error.Code != tt.wantCode {
				t.Errorf("error code: got %s, want %s", res.Error.Code, tt.wantCode)
			}
		})
	}
}
