package apperror_test

import (
	"errors"
	"librigo/internal/domain/apperror"
	"testing"
)

func TestAppError(t *testing.T) {
	baseErr := errors.New("original error")
	code := "TEST_CODE"
	errType := apperror.TypeNotFound

	// 1. New関数のテスト
	err := apperror.New(baseErr, code, errType)

	t.Run("データの保持チェック", func(t *testing.T) {
		if err.Error() != baseErr.Error() {
			t.Errorf("expected error message %s, got %s", baseErr.Error(), err.Error())
		}
		if err.Code != code {
			t.Errorf("expected code %s, got %s", code, err.Code)
		}
		if err.ErrType != errType {
			t.Errorf("expected type %s, got %s", errType, err.ErrType)
		}
	})

	t.Run("Unwrapのテスト", func(t *testing.T) {
		// errors.Is や errors.As が正しく機能するか確認
		if !errors.Is(err, baseErr) {
			t.Error("failed to unwrap: errors.Is(err, baseErr) should be true")
		}

		var target *apperror.AppError
		if !errors.As(err, &target) {
			t.Error("failed to cast: errors.As(err, &target) should be true")
		}
	})
}
