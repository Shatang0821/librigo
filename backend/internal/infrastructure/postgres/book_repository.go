package postgres

import (
	"context"
	"database/sql"
	"errors"
	bookdomain "librigo/internal/domain/book"

	"github.com/jackc/pgx/v5/pgconn"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) bookdomain.BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Save(ctx context.Context, book *bookdomain.Book) error {
	query := `INSERT INTO books (id, title, price, isbn) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query,
		book.ID().String(),
		book.Title().String(),
		book.Price().Int(),
		book.ISBN().String(),
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return bookdomain.ErrDuplicateBook
			}
		}
		return err
	}
	return nil
}

func (r *BookRepository) FindByID(ctx context.Context, id bookdomain.BookID) (*bookdomain.Book, error) {
	query := `SELECT id, title, price, isbn FROM books WHERE id = $1`

	var (
		resID    string
		resTitle string
		resPrice int
		resISBN  string
	)

	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(&resID, &resTitle, &resPrice, &resISBN)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, bookdomain.ErrBookNotFound
		}
		return nil, err
	}

	// DBからの取得なので NewBook でバリデーションをかけつつ生成
	return bookdomain.NewBook(resID, resTitle, resPrice, resISBN)
}

func (r *BookRepository) FindAll(ctx context.Context) ([]*bookdomain.Book, error) {
	query := `SELECT id, title, price, isbn FROM books`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*bookdomain.Book
	for rows.Next() {
		var (
			resID    string
			resTitle string
			resPrice int
			resISBN  string
		)
		if err := rows.Scan(&resID, &resTitle, &resPrice, &resISBN); err != nil {
			return nil, err
		}

		book, err := bookdomain.NewBook(resID, resTitle, resPrice, resISBN)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
