package book

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var (
	// ドメインエラー
	ErrInvalidBookTitle = apperror.New(errors.New("invalid book title: cannot be empty"), "INVALID_BOOK_TITLE", apperror.TypeInvalid)
	ErrInvalidBookPrice = apperror.New(errors.New("invalid book price: cannot be negative"), "INVALID_BOOK_PRICE", apperror.TypeInvalid)
	ErrInvalidBookISBN  = apperror.New(errors.New("invalid book isbn: cannot be empty"), "INVALID_BOOK_ISBN", apperror.TypeInvalid)
	// レポジトリエラー
	ErrBookNotFound  = apperror.New(errors.New("book not found"), "BOOK_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateBook = apperror.New(errors.New("book already exists"), "BOOK_DUPLICATE", apperror.TypeConflict)
)
