package user

import (
	"net/mail"
	"strings"
)

func NewEmail(v string) (Email, error) {
	trimmed := strings.TrimSpace(v)

	addr, err := mail.ParseAddress(trimmed)
	if err != nil {
		return "", ErrInvalidUserEmail
	}
	return Email(addr.Address), nil
}

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
)

type User struct {
	ID           UserID
	Name         string
	Email        Email
	PasswordHash string
	Role         UserRole
}

func NewUser(id UserID, name string, email Email, passwordHash string, role UserRole) (*User, error) {

	// ロールのデフォルト値設定
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
		return ErrWeakUserPassword
	}
	return nil
}
