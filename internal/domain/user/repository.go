package user

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id UserID) (*User, error)
}

type PassWordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
