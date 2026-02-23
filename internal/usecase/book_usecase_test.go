package usecase_test

import (
	"context"
	"librigo/internal/domain/book"
	"librigo/internal/usecase"
	"testing"
)

// --- モックの定義 ---

// MockBookRepository は domain.BookRepository を実装した偽物です
type MockBookRepository struct {
	SaveFunc func(ctx context.Context, b *book.Book) error
	// 必要に応じて FindAllFunc なども追加
}

func (m *MockBookRepository) Save(ctx context.Context, b *book.Book) error {
	return m.SaveFunc(ctx, b)
}
func (m *MockBookRepository) FindAll(ctx context.Context) ([]*book.Book, error) { return nil, nil }
func (m *MockBookRepository) FindByID(ctx context.Context, id book.BookID) (*book.Book, error) {
	return nil, nil
}

// MockIDGenerator は domain.IDGenerator を実装した偽物です
type MockIDGenerator struct {
	ID string
}

func (m *MockIDGenerator) Generate() book.BookID {
	return book.BookID(m.ID)
}

// --- テスト本体 ---

func TestCreateBook(t *testing.T) {
	ctx := context.Background()
	fixedID := "test-uuid"

	// 1. モックの準備
	mockRepo := &MockBookRepository{
		SaveFunc: func(ctx context.Context, b *book.Book) error {
			// リポジトリに渡された値が正しいか検証
			if string(b.ID) != fixedID {
				t.Errorf("expected ID %s, got %s", fixedID, b.ID)
			}
			return nil
		},
	}
	mockIDGen := &MockIDGenerator{ID: fixedID}

	// 2. UseCaseの初期化
	uc := usecase.NewBookUseCase(mockRepo, mockIDGen)

	// 3. 実行
	input := usecase.CreateBookInput{
		Title: "テスト書籍",
		Price: 2000,
		ISBN:  "978-4-0000-0000-0",
	}
	output, err := uc.CreateBook(ctx, input)

	// 4. 検証
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output.ID != fixedID {
		t.Errorf("expected output ID %s, got %s", fixedID, output.ID)
	}
}
