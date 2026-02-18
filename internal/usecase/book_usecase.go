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
	book, err := bookdomain.NewBook(i.idGen.Generate(), input.Title, input.Price, input.ISBN)
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

func (i *bookInteractor) GetAllBooks(ctx context.Context) ([]*BookOutput, error) {
	books, err := i.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []*BookOutput
	for _, b := range books {
		outputs = append(outputs, &BookOutput{
			ID:    string(b.ID),
			Title: b.Title,
			Price: b.Price,
			ISBN:  b.ISBN,
		})
	}
	return outputs, nil
}

func (i *bookInteractor) GetBookByID(ctx context.Context, id string) (*BookOutput, error) {
	book, err := i.repo.FindByID(ctx, bookdomain.BookID(id))
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, bookdomain.ErrBookNotFound
	}

	return &BookOutput{
		ID:    string(book.ID),
		Title: book.Title,
		Price: book.Price,
		ISBN:  book.ISBN,
	}, nil
}
