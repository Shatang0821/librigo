package user

import (
	"context"
	"time"
)

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email UserEmail) (*User, error)
	FindByID(ctx context.Context, id UserID) (*User, error)
}

type PassWordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type IDGenerator interface {
	Generate() UserID
}

// トークンに埋め込む情報を保持する構造体
type UserClaims struct {
	UserID UserID
	Role   UserRole
}

type TokenGenerator interface {
	// User情報をもとにJWTトークンを生成するメソッド
	Generate(u *User, duration time.Duration) (string, error)
	// トークンを解析してUserClaimsを取得するメソッド
	Parse(tokenString string) (*UserClaims, error)
}
