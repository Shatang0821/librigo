package apperror

import (
	"errors"
	"fmt"
	"testing"
)

func TestAppError_MapDriven(t *testing.T) {
	// 共通のテストデータ
	defA := New("ERR_A", TypeInvalid)
	defB := New("ERR_B", TypeConflict)
	cause := errors.New("original cause")

	tests := map[string]struct {
		setup    func() error
		target   error
		wantIs   bool
		wantStr  string
		wantType ErrorType
	}{
		"正常系: AppErrorの情報取得": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			wantStr:  fmt.Sprintf("ERR_A: %v", cause),
			wantType: TypeInvalid,
		},
		"判定系: 同一CodeのAppError比較": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			target: defA.Wrap(errors.New("different cause")), // Codeが同じなら一致
			wantIs: true,
		},
		"判定系: 異なるCodeのAppError比較": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			target: defB.Wrap(cause),
			wantIs: false,
		},
		"判定系: ErrorDefとの比較": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			target: defA,
			wantIs: true,
		},
		"判定系: 全く別のエラーとの比較": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			target: errors.New("standard error"),
			wantIs: false,
		},
		"正常系: Unwrapによる原因の抽出": {
			setup: func() error {
				return defA.Wrap(cause)
			},
			// 検証は Loop 内の errors.Unwrap で行う
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.setup()

			// 1. errors.Is の検証
			if tt.target != nil {
				if got := errors.Is(err, tt.target); got != tt.wantIs {
					t.Errorf("errors.Is() = %v, want %v", got, tt.wantIs)
				}
			}

			// 2. メソッド・文字列の検証
			if appErr, ok := err.(*AppError); ok {
				if tt.wantStr != "" && appErr.Error() != tt.wantStr {
					t.Errorf("Error() = %q, want %q", appErr.Error(), tt.wantStr)
				}
				if tt.wantType != "" && appErr.Type() != tt.wantType {
					t.Errorf("Type() = %v, want %v", appErr.Type(), tt.wantType)
				}
			}

			// 3. Unwrap の検証
			if name == "正常系: Unwrapによる原因の抽出" {
				if got := errors.Unwrap(err); got != cause {
					t.Errorf("Unwrap() = %v, want %v", got, cause)
				}
			}
		})
	}
}
