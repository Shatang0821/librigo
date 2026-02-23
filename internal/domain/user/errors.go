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
)
