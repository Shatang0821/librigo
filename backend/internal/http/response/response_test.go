package response

import (
	"encoding/json"
	"errors"
	apperror "librigo/internal/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse_MapDriven(t *testing.T) {
	// テスト用の型
	type testData struct {
		Message string `json:"message"`
	}

	// モック用のエラー（HTTPStatus フィールドがあると仮定）
	errBadRequest := &apperror.AppError{
		Code:       "BAD_REQUEST",
		Message:    "bad request error",
		HTTPStatus: http.StatusBadRequest,
	}

	tests := map[string]struct {
		// 入力
		inputData interface{}
		inputErr  error
		isError   bool // HandleError を呼ぶかどうか

		// 期待値
		wantStatus int
		wantBody   string // 簡易的な部分一致チェック用
	}{
		"成功: 構造体データを正常に返す": {
			inputData:  testData{Message: "hello"},
			isError:    false,
			wantStatus: http.StatusOK,
			wantBody:   `{"data":{"message":"hello"},"error":null}`,
		},
		"成功: nilデータを正常に返す": {
			inputData:  nil,
			isError:    false,
			wantStatus: http.StatusOK,
			wantBody:   `{"data":null,"error":null}`,
		},
		"失敗: 定義済みAppErrorのハンドリング": {
			inputErr:   errBadRequest,
			isError:    true,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":{"Def":{"Code":"BAD_REQUEST"`,
		},
		"失敗: 未定義エラーはInternalに変換": {
			inputErr:   errors.New("raw error"),
			isError:    true,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// 1. httptest.NewRecorder で書き込み先を準備
			w := httptest.NewRecorder()

			// 2. 対象関数の実行
			if tt.isError {
				HandleError(w, tt.inputErr)
			} else {
				WriteSuccess(w, tt.inputData)
			}

			// 3. ステータスコードの検証
			if w.Code != tt.wantStatus {
				t.Errorf("Status Code = %d, want %d", w.Code, tt.wantStatus)
			}

			// 4. ボディの検証（JSONとしてパースできるか）
			var res APIResponse
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			// 5. Content-Type の検証
			if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
				t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
			}
		})
	}
}
