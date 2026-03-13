package book_test

import (
	"errors"
	"librigo/internal/domain/book"
	"testing"
)

func TestNewBook(t *testing.T) {
	// 正常系用の共通データ
	vID, _ := book.NewBookID("550e8400-e29b-41d4-a716-446655440000")
	vTitle, _ := book.NewBookTitle("テスト駆動開発")
	vPrice, _ := book.NewBookPrice(3000)
	vISBN, _ := book.NewBookISBN("978-4-0000-0000-0")

	tests := map[string]struct {
		id      book.BookID
		title   book.BookTitle
		price   book.BookPrice
		isbn    book.BookISBN
		wantErr error
	}{
		"正常系: 有効な入力": {
			id:      vID,
			title:   vTitle,
			price:   vPrice,
			isbn:    vISBN,
			wantErr: nil,
		},
		"異常系: IDが空": {
			id:      book.BookID{},
			title:   vTitle,
			price:   vPrice,
			isbn:    vISBN,
			wantErr: book.ErrInvalidBook,
		},
		"異常系: タイトルが空": {
			id:      vID,
			title:   book.BookTitle{},
			price:   vPrice,
			isbn:    vISBN,
			wantErr: book.ErrInvalidBook,
		},
		"異常系: ISBNが空": {
			id:      vID,
			title:   vTitle,
			price:   vPrice,
			isbn:    book.BookISBN{},
			wantErr: book.ErrInvalidBook,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := book.NewBook(tt.id, tt.title, tt.price, tt.isbn)

			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			} else {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("errors.Is() = false, want true\n got: %v\n want: %v", err, tt.wantErr)
				}
			}

			if tt.wantErr == nil && got != nil {
				if got.ID() != tt.id {
					t.Errorf("expected ID %v, got %v", tt.id, got.ID())
				}
				if got.Title() != tt.title {
					t.Errorf("expected Title %v, got %v", tt.title, got.Title())
				}
				if got.Price() != tt.price {
					t.Errorf("expected Price %v, got %v", tt.price, got.Price())
				}
				if got.ISBN() != tt.isbn {
					t.Errorf("expected ISBN %v, got %v", tt.isbn, got.ISBN())
				}
			}
		})
	}
}
