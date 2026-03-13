package book

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var (
	ErrInvalidBookTitle = apperror.New("INVALID_BOOK_TITLE", apperror.TypeInvalid)
	ErrInvalidBookPrice = apperror.New("INVALID_BOOK_PRICE", apperror.TypeInvalid)
	ErrInvalidBookISBN  = apperror.New("INVALID_BOOK_ISBN", apperror.TypeInvalid)
)

// BookID は本の一意識別子です
type BookID struct{ value string }

func NewBookID(v string) (BookID, error) {
	if v == "" {
		return BookID{}, ErrInvalidBookISBN.Wrap(errors.New("book ID can not be empty"))
	}
	return BookID{value: v}, nil
}

func (id BookID) String() string {
	return id.value
}

// BookTitle は本のタイトルです
type BookTitle struct{ value string }

func NewBookTitle(v string) (BookTitle, error) {
	if v == "" {
		return BookTitle{}, ErrInvalidBookTitle.Wrap(errors.New("book title can not be empty"))
	}
	return BookTitle{value: v}, nil
}

func (t BookTitle) String() string {
	return t.value
}

// BookPrice は本の価格です
type BookPrice struct{ value int }

func NewBookPrice(v int) (BookPrice, error) {
	if v < 0 {
		return BookPrice{}, ErrInvalidBookPrice.Wrap(errors.New("book price can not be negative"))
	}
	return BookPrice{value: v}, nil
}

func (p BookPrice) Int() int {
	return p.value
}

// BookISBN は本のISBNコードです
type BookISBN struct{ value string }

func NewBookISBN(v string) (BookISBN, error) {
	if v == "" {
		return BookISBN{}, ErrInvalidBookISBN.Wrap(errors.New("book ISBN can not be empty"))
	}
	return BookISBN{value: v}, nil
}

func (i BookISBN) String() string {
	return i.value
}
