package user

import (
	"errors"
	"testing"
)

func TestNewUser_MapDriven(t *testing.T) {
	// 正常系用の共通データ
	vID := "user-001"
	vName := "Gopher"
	vEmail := "gopher@example.com"
	vPass := "hashed-string"
	vRole := "admin"

	tests := map[string]struct {
		id      string
		name    string
		email   string
		pw      string
		role    string
		wantErr error
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
			id:      "",
			name:    vName,
			email:   vEmail,
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUserID,
		},
		"異常系: 名前が空": {
			id:      vID,
			name:    "",
			email:   vEmail,
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUserName,
		},
		"異常系: メールが無効": {
			id:      vID,
			name:    vName,
			email:   "invalid-email",
			pw:      vPass,
			role:    vRole,
			wantErr: ErrInvalidUserEmail,
		},
		"異常系: パスワードハッシュが空": {
			id:      vID,
			name:    vName,
			email:   vEmail,
			pw:      "",
			role:    vRole,
			wantErr: ErrInvalidUserPassword,
		},
		"異常系: ロールが無効": {
			id:      vID,
			name:    vName,
			email:   vEmail,
			pw:      vPass,
			role:    "invalid-role",
			wantErr: ErrInvalidUserRole,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewUser(tt.id, tt.name, tt.email, tt.pw, tt.role)

			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			} else {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("errors.Is() = false, want true\n got: %v\n want: %v", err, tt.wantErr)
				}
			}

			if tt.wantErr == nil && got != nil {
				if got.ID().String() != tt.id {
					t.Errorf("ID() = %v, want %v", got.ID().String(), tt.id)
				}
				if got.Name().String() != tt.name {
					t.Errorf("Name() = %v, want %v", got.Name().String(), tt.name)
				}
				if got.Email().String() != tt.email {
					t.Errorf("Email() = %v, want %v", got.Email().String(), tt.email)
				}
				if got.PasswordHash().String() != tt.pw {
					t.Errorf("PasswordHash() = %v, want %v", got.PasswordHash().String(), tt.pw)
				}
				if got.Role().String() != tt.role {
					t.Errorf("Role() = %v, want %v", got.Role().String(), tt.role)
				}
			}
		})
	}
}
