package apperror

import (
	"errors"
	"fmt"
	"testing"
)

func TestAppError(t *testing.T) {
	// テスト用の定義と元となるエラー
	defA := New("ERR_A", TypeInvalid)
	defB := New("ERR_B", TypeConflict)
	baseErr := errors.New("original error message")

	t.Run("Getter methods and Error()", func(t *testing.T) {
		appErr := defA.WithError(baseErr)

		if appErr.GetCode() != "ERR_A" {
			t.Errorf("expected code ERR_A, got %s", appErr.GetCode())
		}
		if appErr.GetType() != TypeInvalid {
			t.Errorf("expected type %s, got %s", TypeInvalid, appErr.GetType())
		}
		if appErr.Error() != "original error message" {
			t.Errorf("expected message 'original error message', got %s", appErr.Error())
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		appErr := defA.WithError(baseErr)
		unwrapped := errors.Unwrap(appErr)

		if unwrapped != baseErr {
			t.Errorf("unwrapped error does not match original")
		}
	})

	t.Run("Is function (errors.Is compliance)", func(t *testing.T) {
		tests := []struct {
			name   string
			err    error
			target error
			want   bool
		}{
			{
				name:   "一致する場合",
				err:    defA.WithError(baseErr),
				target: defA,
				want:   true,
			},
			{
				name:   "別のErrorDefと比較する場合",
				err:    defA.WithError(baseErr),
				target: defB,
				want:   false,
			},
			{
				name:   "全く別のエラー型と比較する場合",
				err:    defA.WithError(baseErr),
				target: errors.New("other"),
				want:   false,
			},
			{
				name:   "ErrorDef単体での比較 (ErrorDef.Error()の確認)",
				err:    defA,
				target: defA,
				want:   true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Go標準の errors.Is を使用してテスト
				if got := errors.Is(tt.err, tt.target); got != tt.want {
					t.Errorf("errors.Is() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestGlobalErrors(t *testing.T) {
	t.Run("ErrInvalidJSON", func(t *testing.T) {
		err := ErrInvalidJSON.WithError(fmt.Errorf("syntax error"))
		if !errors.Is(err, ErrInvalidJSON) {
			t.Error("should match ErrInvalidJSON")
		}
		if err.GetType() != TypeInvalid {
			t.Error("type should be TypeInvalid")
		}
	})
}
