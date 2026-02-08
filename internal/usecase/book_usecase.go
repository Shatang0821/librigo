package usecase

import (
	"context"
	"librigo/internal/domain"
)

// CreateBoookInput は書籍登録に必要な入力データです
type CreateBookInput struct {
	Title string
	Price int
	ISBN  string
}

// CreateBookOutput は書籍登録の出力データです
type CreateBookOutput struct {
	ID    string
	Title string
}

// BookUseCase は書籍に関するユースケースを定義するインターフェースです
type BookUseCase interface {
	CreateBook(ctx context.Context, input CreateBookInput) (*CreateBookOutput, error)
}

// bookInteractor は BookUseCase の実装です
type bookInteractor struct {
	repo domain.BookRepository
}

func NewBookUseCase(repo domain.BookRepository) BookUseCase {
	return &bookInteractor{repo: repo}
}

func (i *bookInteractor) CreateBook(ctx context.Context, input CreateBookInput) (*CreateBookOutput, error) {
	// ドメインモデルの生成
	// IDの生成ロジックは本来どこかで定義しますが、一旦仮のIDを渡します
	book, err := domain.NewBook(domain.BookID("temp-id"), input.Title, input.Price, input.ISBN)
	if err != nil {
		return nil, err
	}

	// リポジトリを介した保存
	if err := i.repo.Save(ctx, book); err != nil {
		return nil, err
	}

	// 出力データの生成
	return &CreateBookOutput{
		ID:    string(book.ID),
		Title: book.Title,
	}, nil
}
