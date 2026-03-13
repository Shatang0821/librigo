package user

import (
	"errors"
	"testing"
)

func TestValueObjects_MapDriven(t *testing.T) {
	tests := map[string]struct {
		run     func() (any, error)
		wantVal string
		wantErr error
	}{
		// --- UserID / UserName ---
		"正常: 有効なUserID": {
			run:     func() (any, error) { return NewUserID("id-123") },
			wantVal: "id-123",
		},
		"異常: IDが空": {
			run:     func() (any, error) { return NewUserID("") },
			wantErr: ErrInvalidUserID.Wrap(nil),
		},

		// --- UserEmail ---
		"正常: 有効なEmail(トリム)": {
			run:     func() (any, error) { return NewEmail("  test@example.com  ") },
			wantVal: "test@example.com",
		},
		"異常: Emailの形式不正": {
			run:     func() (any, error) { return NewEmail("not-an-email") },
			wantErr: ErrInvalidUserEmail.Wrap(nil),
		},

		// --- UserRawPassword (複雑なバリデーション) ---
		"正常: パスワード要件を満たす": {
			run:     func() (any, error) { return NewUserRawPassword("Valid123") },
			wantVal: "Valid123",
		},
		"異常: パスワードが短い": {
			run:     func() (any, error) { return NewUserRawPassword("V12") },
			wantErr: ErrWeakUserPassword.Wrap(nil),
		},
		"異常: パスワードに数字がない": {
			run:     func() (any, error) { return NewUserRawPassword("NoDigits") },
			wantErr: ErrWeakUserPassword.Wrap(nil),
		},

		// --- UserRole ---
		"正常: adminロールの取得": {
			run:     func() (any, error) { return NewUserRole("admin") },
			wantVal: "admin",
		},
		"異常: 未定義のロール": {
			run:     func() (any, error) { return NewUserRole("superman") },
			wantErr: ErrInvalidUserRole.Wrap(nil),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.run()

			// エラー検証
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("errors.Is() = false\n got: %v\n want: %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// 値の検証 (Stringerインターフェースを利用)
			if s, ok := got.(interface{ String() string }); ok {
				if s.String() != tt.wantVal {
					t.Errorf("got value %q, want %q", s.String(), tt.wantVal)
				}
			}
		})
	}
}

func TestUserRole_Methods(t *testing.T) {
	t.Run("IsAdminの判定", func(t *testing.T) {
		admin, _ := NewUserRole("admin")
		member, _ := NewUserRole("member")

		if !admin.IsAdmin() {
			t.Error("admin.IsAdmin() should be true")
		}
		if member.IsAdmin() {
			t.Error("member.IsAdmin() should be false")
		}
	})
}
