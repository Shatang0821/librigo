package book

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var (
	ErrBookNotFound  = apperror.New(errors.New("book not found"), "BOOK_NOT_FOUND", apperror.TypeNotFound)
	ErrDuplicateBook = apperror.New(errors.New("book already exists"), "BOOK_DUPLICATE", apperror.TypeConflict)
)
