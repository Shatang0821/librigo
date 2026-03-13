package book

import (
	"errors"
	"librigo/internal/domain/apperror"
)

var ErrInvalidBook = apperror.New("INVALID_BOOK", apperror.TypeInvalid)

type Book struct {
	id    BookID
	title BookTitle
	price BookPrice
	isbn  BookISBN
}

// Bookのコンストラクタ
func NewBook(id BookID, title BookTitle, price BookPrice, isbn BookISBN) (*Book, error) {
	if id.value == "" || title.value == "" || isbn.value == "" {
		return nil, ErrInvalidBook.Wrap(errors.New("Book fields cannot be empty"))
	}

	return &Book{
		id:    id,
		title: title,
		price: price,
		isbn:  isbn,
	}, nil
}

// 外部から値を取得するための Getter メソッド群
func (b *Book) ID() BookID       { return b.id }
func (b *Book) Title() BookTitle { return b.title }
func (b *Book) Price() BookPrice { return b.price }
func (b *Book) ISBN() BookISBN   { return b.isbn }
