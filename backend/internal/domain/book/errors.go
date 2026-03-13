package book

import (
	"librigo/internal/domain/apperror"
)

var (
	// ドメインエラー

	// レポジトリエラー
	ErrBookNotFound  = apperror.New("BOOK_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateBook = apperror.New("BOOK_DUPLICATE", apperror.TypeConflict)
)
