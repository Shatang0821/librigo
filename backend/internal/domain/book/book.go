package book

type Book struct {
	id    BookID
	title BookTitle
	price BookPrice
	isbn  BookISBN
}

// NewBook は、プリミティブな値から Book エンティティを生成します。
// 内部で各 Value Object のバリデーションが実行されます。
func NewBook(id string, title string, price int, isbn string) (*Book, error) {
	voID, err := NewBookID(id)
	if err != nil {
		return nil, err
	}

	voTitle, err := NewBookTitle(title)
	if err != nil {
		return nil, err
	}

	voPrice, err := NewBookPrice(price)
	if err != nil {
		return nil, err
	}

	voISBN, err := NewBookISBN(isbn)
	if err != nil {
		return nil, err
	}

	return &Book{
		id:    voID,
		title: voTitle,
		price: voPrice,
		isbn:  voISBN,
	}, nil
}

// 外部から値を取得するための Getter メソッド群
func (b *Book) ID() BookID       { return b.id }
func (b *Book) Title() BookTitle { return b.title }
func (b *Book) Price() BookPrice { return b.price }
func (b *Book) ISBN() BookISBN   { return b.isbn }
