package user

import (
	"errors"
	"librigo/internal/domain/apperror"
)

// ドメインエラー
var (
	ErrInvalidUserID       = apperror.New("INVALID_USER_ID", apperror.TypeInvalid)
	ErrInvalidUserName     = apperror.New("INVALID_USER_NAME", apperror.TypeInvalid)
	ErrInvalidUserEmail    = apperror.New("INVALID_USER_EMAIL", apperror.TypeInvalid)
	ErrWeakUserPassword    = apperror.New("WEAK_USER_PASSWORD", apperror.TypeInvalid)
	ErrInvalidUserPassword = apperror.New("INVALID_USER_PASSWORD", apperror.TypeInvalid)
)

// UserID はユーザーの一意識別子です
type UserID struct{ value string }

func NewUserID(v string) (UserID, error) {
	if v == "" {
		return UserID{}, ErrInvalidUserID.WithError(errors.New("User ID can not be empty"))
	}
	return UserID{value: v}, nil
}

type UserName struct{ value string }

func NewUserName(v string) (UserName, error) {
	if v == "" {
		return UserName{}, ErrInvalidUserName.WithError(errors.New("User name can not be empty"))
	}
	return UserName{value: v}, nil
}

// UserRole はユーザーの権限を定義します
type UserRole string

// Email はユーザーのメールアドレスを表す型です
type Email string
