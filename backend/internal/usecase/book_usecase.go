package usecase

import (
	"context"
	bookdomain "librigo/internal/domain/book"
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

// BookOutput は書籍の出力データです
type BookOutput struct {
	ID    string
	Title string
	Price int
	ISBN  string
}

// BookUseCase は書籍に関するユースケースを定義するインターフェースです
type BookUseCase interface {
	CreateBook(ctx context.Context, input CreateBookInput) (*CreateBookOutput, error)
	GetAllBooks(ctx context.Context) ([]*BookOutput, error)
	GetBookByID(ctx context.Context, id string) (*BookOutput, error)
}

// bookInteractor は BookUseCase の実装です
type bookInteractor struct {
	repo  bookdomain.BookRepository
	idGen bookdomain.IDGenerator
}

func NewBookUseCase(repo bookdomain.BookRepository, idGen bookdomain.IDGenerator) BookUseCase {
	return &bookInteractor{repo: repo, idGen: idGen}
}

func (i *bookInteractor) CreateBook(ctx context.Context, input CreateBookInput) (*CreateBookOutput, error) {
	// ドメインモデルの生成
	// NewBook コンストラクタ内で Value Object への変換とバリデーションが行われます
	book, err := bookdomain.NewBook(
		i.idGen.Generate().String(),
		input.Title,
		input.Price,
		input.ISBN,
	)
	if err != nil {
		return nil, err
	}

	// リポジトリを介した保存
	if err := i.repo.Save(ctx, book); err != nil {
		return nil, err
	}

	// 出力データの生成
	return &CreateBookOutput{
		ID:    book.ID().String(),
		Title: book.Title().String(),
	}, nil
}

func (i *bookInteractor) GetAllBooks(ctx context.Context) ([]*BookOutput, error) {
	books, err := i.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []*BookOutput
	for _, b := range books {
		outputs = append(outputs, &BookOutput{
			ID:    string(b.ID().String()),
			Title: b.Title().String(),
			Price: b.Price().Int(),
			ISBN:  b.ISBN().String(),
		})
	}
	return outputs, nil
}

func (i *bookInteractor) GetBookByID(ctx context.Context, id string) (*BookOutput, error) {
	bookId, err := bookdomain.NewBookID(id)
	if err != nil {
		return nil, err
	}
	book, err := i.repo.FindByID(ctx, bookId)
	if err != nil {
		return nil, err
	}

	return &BookOutput{
		ID:    string(book.ID().String()),
		Title: book.Title().String(),
		Price: book.Price().Int(),
		ISBN:  book.ISBN().String(),
	}, nil
}
