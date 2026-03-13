package user

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var (
	ErrInvalidUser = apperror.New("INVALID_USER", apperror.TypeInvalid)
)

type User struct {
	id           UserID
	name         UserName
	email        UserEmail
	passwordHash UserHashedPassword
	role         UserRole
}

func NewUser(id UserID, name UserName, email UserEmail, passwordHash UserHashedPassword, role UserRole) (*User, error) {
	if id.value == "" || name.value == "" || email.value == "" || passwordHash.value == "" || role.value == "" {
		return nil, ErrInvalidUser.Wrap(errors.New("User fields cannot be empty"))
	}

	return &User{
		id:           id,
		name:         name,
		email:        email,
		passwordHash: passwordHash,
		role:         role,
	}, nil
}

// 外部から値を取得するための Getter メソッド群
func (u *User) ID() UserID                       { return u.id }
func (u *User) Name() UserName                   { return u.name }
func (u *User) Email() UserEmail                 { return u.email }
func (u *User) Role() UserRole                   { return u.role }
func (u *User) PasswordHash() UserHashedPassword { return u.passwordHash }
