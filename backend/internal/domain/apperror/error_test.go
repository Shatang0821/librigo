package apperror

import (
	"errors"
	"testing"
)

func TestAppError_TableDriven(t *testing.T) {
	// 固定のテスト用定義
	testDef := New("TEST_ERR", TypeInvalid)
	otherDef := New("OTHER_ERR", TypeConflict)
	baseErr := errors.New("original message")

	// map形式でのテストケース定義
	tests := map[string]struct {
		// 入力とセットアップ
		setup  func() error
		target error // errors.Is で比較する対象

		// 期待値
		wantIs     bool
		expectCode string
		expectType ErrorType
	}{
		"正常系: AppErrorの生成と情報取得": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target:     nil, // Isのテストは行わない
			wantIs:     false,
			expectCode: "TEST_ERR",
			expectType: TypeInvalid,
		},
		"判定系: errors.Isでの一致(同一インスタンス)": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target: testDef,
			wantIs: true,
		},
		"判定系: errors.Isでの一致(別インスタンス/同Code)": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target: &ErrorDef{Code: "TEST_ERR", ErrType: TypeInvalid},
			wantIs: true,
		},
		"判定系: errors.Isでの不一致(別Code)": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target: otherDef,
			wantIs: false,
		},
		"判定系: errors.Isでの不一致(型違い)": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target: errors.New("TEST_ERR"),
			wantIs: false,
		},
		"正常系: Unwrapによる元エラーの取得": {
			setup: func() error {
				return testDef.WithError(baseErr)
			},
			target: nil,
			wantIs: false,
			// Unwrapの検証は別途行うが、期待値として定義しても良い
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// テスト対象のエラーを生成
			err := tt.setup()

			// 1. errors.Is の検証
			if tt.target != nil {
				if got := errors.Is(err, tt.target); got != tt.wantIs {
					t.Errorf("errors.Is() = %v, want %v", got, tt.wantIs)
				}
			}

			// 2. Getter関数の検証（AppError型にキャスト可能な場合のみ）
			if appErr, ok := err.(*AppError); ok {
				if tt.expectCode != "" && appErr.GetCode() != tt.expectCode {
					t.Errorf("GetCode() = %v, want %v", appErr.GetCode(), tt.expectCode)
				}
				if tt.expectType != "" && appErr.GetType() != tt.expectType {
					t.Errorf("GetType() = %v, want %v", appErr.GetType(), tt.expectType)
				}
			}

			// 3. Unwrapの検証（特定のケースのみ）
			if name == "正常系: Unwrapによる元エラーの取得" {
				unwrapped := errors.Unwrap(err)
				if unwrapped != baseErr {
					t.Errorf("Unwrap() failed to return baseErr")
				}
			}
		})
	}
}
