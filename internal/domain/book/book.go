package book

// カスタム型定義
type BookID string

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
