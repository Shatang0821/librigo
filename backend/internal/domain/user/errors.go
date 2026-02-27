package user

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var (
	// ドメインエラー
	ErrInvalidUserEmail = apperror.New(errors.New("Invalid email address format"), "INVALID_EMAIL", apperror.TypeInvalid)
	ErrInvalidUserName  = apperror.New(errors.New("User name cannot be empty"), "INVALID_USER_NAME", apperror.TypeInvalid)
	ErrWeakPassword     = apperror.New(errors.New("Password is too weak"), "WEAK_PASSWORD", apperror.TypeInvalid)
	// レポジトリエラー
	ErrUserNotFound  = apperror.New(errors.New("User not found"), "USER_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateUser = apperror.New(errors.New("User already exists"), "USER_DUPLICATE", apperror.TypeConflict)
)
