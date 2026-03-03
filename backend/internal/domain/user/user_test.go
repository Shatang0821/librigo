package user_test

import (
	"errors"
	"librigo/internal/domain/user"
	"testing"
)

func TestNewUser(t *testing.T) {
	//更新必要
	tests := map[string]struct {
		id           string
		name         string
		email        string
		passwordHash string
		role         user.UserRole
		wantErr      error
		wantRole     user.UserRole // デフォルト値の確認用
	}{
		"正常系: 有効なユーザー": {
			id:           "user-1",
			name:         "山田太郎",
			email:        "taro@example.com",
			passwordHash: "hashed_password",
			role:         user.RoleAdmin,
			wantErr:      nil,
			wantRole:     user.RoleAdmin,
		},
		"正常系: ロール未指定時はMemberになる": {
			id:           "user-2",
			name:         "佐藤次郎",
			email:        "jiro@example.com",
			passwordHash: "hashed_password",
			role:         "", // 未指定
			wantErr:      nil,
			wantRole:     user.RoleMember,
		},
		"異常系: 名前が空": {
			id:           "user-3",
			name:         "",
			email:        "test@example.com",
			passwordHash: "hashed_password",
			wantErr:      user.ErrInvalidUserName,
		},
		"異常系: メールアドレスが空": {
			id:           "user-4",
			name:         "テストユーザー",
			email:        "",
			passwordHash: "hashed_password",
			wantErr:      user.ErrInvalidUserEmail,
		},
		"異常系: パスワードが空": {
			id:           "user-5",
			name:         "テストユーザー",
			email:        "test@example.com",
			passwordHash: "",
			wantErr:      user.ErrInvalidUserPassword,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := user.NewUser(
				user.UserID(tt.id),
				tt.name,
				user.Email(tt.email),
				tt.passwordHash,
				tt.role,
			)

			// エラーの検証
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 成功時のデータ検証
			if tt.wantErr == nil {
				if got.Role != tt.wantRole {
					t.Errorf("expected role %s, got %s", tt.wantRole, got.Role)
				}
				if got.Email != user.Email(tt.email) {
					t.Errorf("expected email %s, got %s", tt.email, got.Email)
				}
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := map[string]struct {
		password string
		wantErr  error
	}{
		"正常系: 8文字ちょうど": {password: "12345678", wantErr: nil},
		"正常系: 長いパスワード": {password: "super-strong-password-123", wantErr: nil},
		"異常系: 短すぎる":    {password: "1234567", wantErr: user.ErrWeakUserPassword},
		"異常系: 空文字":     {password: "", wantErr: user.ErrWeakUserPassword},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := user.ValidatePassword(tt.password)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}
