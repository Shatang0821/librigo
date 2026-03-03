package user

import (
	"librigo/internal/domain/apperror"
)

var (
	// 認証エラー
	ErrInvalidCredentials = apperror.New("INVALID_CREDENTIALS", apperror.TypeUnauthorized)
	ErrUnauthorized       = apperror.New("UNAUTHORIZED", apperror.TypeUnauthenticated)
	// レポジトリエラー
	ErrUserNotFound  = apperror.New("USER_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateUser = apperror.New("USER_DUPLICATE", apperror.TypeConflict)
)
