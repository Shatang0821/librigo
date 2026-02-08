package domain

import "errors"

// カスタム型定義
type BookID string

var (
	ErrInvalidBookTitle = errors.New("invalid book title: cannot be empty")
	ErrInvalidBookPrice = errors.New("invalid book price: cannot be negative")
)

type Book struct {
	ID    BookID
	Title string
	Price int
	ISBN  string
}

// Bookのコンストラクタ
func NewBook(id BookID, title string, price int, isbn string) (*Book, error) {
	if title == "" {
		return nil, ErrInvalidBookTitle
	}
	if price < 0 {
		return nil, ErrInvalidBookPrice
	}
	return &Book{
		ID:    id,
		Title: title,
		Price: price,
		ISBN:  isbn,
	}, nil
}
