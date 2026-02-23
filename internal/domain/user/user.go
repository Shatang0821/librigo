package user

import "net/mail"

// UserID はユーザーの一意識別子です
type UserID string

// UserRole はユーザーの権限を定義します
type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
)

type User struct {
	ID           UserID
	Name         string
	Email        string
	PasswordHash string
	Role         UserRole
}

func NewUser(id UserID, name, email, passwordHash string, role UserRole) (*User, error) {
	if name == "" {
		return nil, ErrInvalidUserName
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrInvalidUserEmail
	}

	// ロールのデフォルト値設定（未指定なら Member）
	if role == "" {
		role = RoleMember
	}

	return &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}, nil
}

// ValidatePassword はパスワードの強度を検証します
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrWeakPassword
	}
	return nil
}
