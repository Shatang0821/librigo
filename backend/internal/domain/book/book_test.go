package book_test

import (
	"errors"
	"librigo/internal/domain/book"
	"testing"
)

func TestNewBook(t *testing.T) {
	// テストケースの定義
	tests := map[string]struct {
		id      string
		title   string
		price   int
		isbn    string
		wantErr error
	}{
		"正常系: 有効な入力": {
			id:      "550e8400-e29b-41d4-a716-446655440000",
			title:   "テスト駆動開発",
			price:   3000,
			isbn:    "978-4-0000-0000-0",
			wantErr: nil,
		},
		"異常系: タイトルが空": {
			id:      "550e8400-e29b-41d4-a716-446655440000",
			title:   "",
			price:   3000,
			isbn:    "978-4-0000-0000-0",
			wantErr: book.ErrInvalidBookTitle,
		},
		"異常系: ISBNが空": {
			id:      "550e8400-e29b-41d4-a716-446655440000",
			title:   "サンプル本",
			price:   3000,
			isbn:    "",
			wantErr: book.ErrInvalidBookISBN,
		},
		"異常系: 価格がマイナス": {
			id:      "550e8400-e29b-41d4-a716-446655440000",
			title:   "サンプル本",
			price:   -1,
			isbn:    "978-4-0000-0000-0",
			wantErr: book.ErrInvalidBookPrice,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// 並列実行
			t.Parallel()

			got, err := book.NewBook(book.BookID(tt.id), tt.title, tt.price, tt.isbn)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == nil {
				if string(got.ID) != tt.id {
					t.Errorf("expected ID %s, got %s", tt.id, got.ID)
				}
				if got.Title != tt.title {
					t.Errorf("expected Title %s, got %s", tt.title, got.Title)
				}
				if got.Price != tt.price {
					t.Errorf("expected Price %d, got %d", tt.price, got.Price)
				}
				if got.ISBN != tt.isbn {
					t.Errorf("expected ISBN %s, got %s", tt.isbn, got.ISBN)
				}
			}
		})
	}

}
