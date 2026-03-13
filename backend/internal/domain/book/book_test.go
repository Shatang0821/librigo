package book_test

import (
	"errors"
	"librigo/internal/domain/book"
	"testing"
)

func TestNewBook(t *testing.T) {
	// 正常系用の共通データ
	validID := "550e8400-e29b-41d4-a716-446655440000"
	validTitle := "テスト駆動開発"
	validPrice := 3000
	validISBN := "978-4-0000-0000-0"

	tests := map[string]struct {
		id      string
		title   string
		price   int
		isbn    string
		wantErr error
	}{
		"正常系: 有効な入力": {
			id:      validID,
			title:   validTitle,
			price:   validPrice,
			isbn:    validISBN,
			wantErr: nil,
		},
		"異常系: IDが空": {
			id:      "",
			title:   validTitle,
			price:   validPrice,
			isbn:    validISBN,
			wantErr: book.ErrInvalidBookID,
		},
		"異常系: タイトルが空": {
			id:      validID,
			title:   "",
			price:   validPrice,
			isbn:    validISBN,
			wantErr: book.ErrInvalidBookTitle,
		},
		"異常系: 価格がマイナス": {
			id:      validID,
			title:   validTitle,
			price:   -1,
			isbn:    validISBN,
			wantErr: book.ErrInvalidBookPrice,
		},
		"異常系: ISBNが空": {
			id:      validID,
			title:   validTitle,
			price:   validPrice,
			isbn:    "",
			wantErr: book.ErrInvalidBookISBN,
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
				if got.ID().String() != tt.id {
					t.Errorf("expected ID %v, got %v", tt.id, got.ID().String())
				}
				if got.Title().String() != tt.title {
					t.Errorf("expected Title %v, got %v", tt.title, got.Title().String())
				}
				if got.Price().Int() != tt.price {
					t.Errorf("expected Price %v, got %v", tt.price, got.Price().Int())
				}
				if got.ISBN().String() != tt.isbn {
					t.Errorf("expected ISBN %v, got %v", tt.isbn, got.ISBN().String())
				}
			}
		})
	}
}
