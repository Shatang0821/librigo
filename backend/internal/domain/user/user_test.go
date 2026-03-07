package user

import (
	"errors"
	"testing"
)

func TestNewUser(t *testing.T) {
	// 有効なテストデータの準備
	validID := UserID{value: "user-123"}
	validName := UserName{value: "Tanaka"}
	validEmail := UserEmail{value: "test@example.com"}
	validPass := UserHashedPassword{value: "$2a$10$..."}
	validRole := UserRole{value: "admin"}

	tests := []struct {
		name    string
		id      UserID
		un      UserName
		email   UserEmail
		pw      UserHashedPassword
		role    UserRole
		wantErr bool
		target  error // errors.Is で判定する対象
	}{
		{
			name:    "【正常系】正しい入力でユーザーが作成される",
			id:      validID,
			un:      validName,
			email:   validEmail,
			pw:      validPass,
			role:    validRole,
			wantErr: false,
		},
		{
			name:    "【異常系】IDが空",
			id:      UserID{value: ""},
			un:      validName,
			email:   validEmail,
			pw:      validPass,
			role:    validRole,
			wantErr: true,
			target:  ErrInvalidUser,
		},
		{
			name:    "【異常系】名前が空",
			id:      validID,
			un:      UserName{value: ""},
			email:   validEmail,
			pw:      validPass,
			role:    validRole,
			wantErr: true,
			target:  ErrInvalidUser,
		},
		// 他のフィールド（email, pw, role）も同様に作成可能
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.id, tt.un, tt.email, tt.pw, tt.role)

			// エラーの有無のチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 独自エラー型（apperror）が正しく返ってきているかのチェック
			if tt.wantErr && tt.target != nil {
				if !errors.Is(err, tt.target) {
					t.Errorf("expected error to be %v, but got %v", tt.target, err)
				}
			}

			// 正常系の場合、Getterで値が正しくセットされているか確認
			if !tt.wantErr {
				if got.ID() != tt.id || got.Name() != tt.un {
					t.Errorf("Getter values do not match input")
				}
			}
		})
	}
}
