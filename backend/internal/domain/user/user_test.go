package user

import (
	"errors"
	"testing"
)

func TestNewUser_MapDriven(t *testing.T) {
	// 正常系用の共通データ
	vID := UserID{value: "user-001"}
	vName := UserName{value: "Gopher"}
	vEmail := UserEmail{value: "gopher@example.com"}
	vPass := UserHashedPassword{value: "hashed-string"}
	vRole := UserRole{value: "admin"}

	tests := map[string]struct {
		id      UserID
		name    UserName
		email   UserEmail
		pw      UserHashedPassword
		role    UserRole
		wantErr error // errors.Is で判定する対象
	}{
		"正常系: 全てのフィールドが有効": {
			id:      vID,
			name:    vName,
			email:   vEmail,
			pw:      vPass,
			role:    vRole,
			wantErr: nil,
		},
		"異常系: IDが空": {
			id:      UserID{value: ""},
			name:    vName,
			email:   vEmail,
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUser.Wrap(nil), // Codeが一致すればOK
		},
		"異常系: 名前が空": {
			id:      vID,
			name:    UserName{value: ""},
			email:   vEmail,
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUser.Wrap(nil),
		},
		"異常系: メールが空": {
			id:      vID,
			name:    vName,
			email:   UserEmail{value: ""},
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUser.Wrap(nil),
		},
		"異常系: パスワードハッシュが空": {
			id:      vID,
			name:    vName,
			email:   vEmail,
			pw:      UserHashedPassword{value: ""},
			role:    vRole,
			wantErr: ErrInvalidUser.Wrap(nil),
		},
		"異常系: ロールが空": {
			id:      vID,
			name:    vName,
			email:   vEmail,
			pw:      vPass,
			role:    UserRole{value: ""},
			wantErr: ErrInvalidUser.Wrap(nil),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewUser(tt.id, tt.name, tt.email, tt.pw, tt.role)

			// 1. エラーの判定
			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			} else {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("errors.Is() = false, want true\n got: %v\n want: %v", err, tt.wantErr)
				}
			}

			// 2. 正常系の場合、Getterが正しく値を返すか検証
			if tt.wantErr == nil && got != nil {
				if got.ID() != tt.id {
					t.Errorf("ID() = %v, want %v", got.ID(), tt.id)
				}
				if got.Name() != tt.name {
					t.Errorf("Name() = %v, want %v", got.Name(), tt.name)
				}
				if got.Email() != tt.email {
					t.Errorf("Email() = %v, want %v", got.Email(), tt.email)
				}
				if got.PasswordHash() != tt.pw {
					t.Errorf("PasswordHash() = %v, want %v", got.PasswordHash(), tt.pw)
				}
				if got.Role() != tt.role {
					t.Errorf("Role() = %v, want %v", got.Role(), tt.role)
				}
			}
		})
	}
}
