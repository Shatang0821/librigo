package book

import (
	"librigo/internal/domain/apperror"
)

var (
	// ドメインエラー
	ErrInvalidBookTitle = apperror.New("INVALID_BOOK_TITLE", apperror.TypeInvalid)
	ErrInvalidBookPrice = apperror.New("INVALID_BOOK_PRICE", apperror.TypeInvalid)
	ErrInvalidBookISBN  = apperror.New("INVALID_BOOK_ISBN", apperror.TypeInvalid)
	// レポジトリエラー
	ErrBookNotFound  = apperror.New("BOOK_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateBook = apperror.New("BOOK_DUPLICATE", apperror.TypeConflict)
)
