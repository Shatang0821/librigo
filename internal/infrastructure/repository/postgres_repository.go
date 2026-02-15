package repository

import (
	"context"
	"database/sql"
	"errors"
	"librigo/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) domain.BookRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(ctx context.Context, book *domain.Book) error {
	query := `INSERT INTO books (id, title, price, isbn) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, book.ID, book.Title, book.Price, book.ISBN)
	if err != nil {
		// PostgreSQLの固有エラーかどうかをチェック
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// "23505" は unique_violation (一意制約違反) のコードです
			if pgErr.Code == "23505" {
				return domain.ErrDuplicateBook
			}
		}
		// それ以外のエラー（接続不良など）はそのまま返す
		return err
	}
	return nil
}

func (r *PostgresRepository) FindByID(ctx context.Context, id domain.BookID) (*domain.Book, error) {
	query := `SELECT id, title, price, isbn FROM books WHERE id = $1`

	var book domain.Book

	err := r.db.QueryRowContext(ctx, query, id).Scan(&book.ID, &book.Title, &book.Price, &book.ISBN)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound // 見つからない場合は domain.ErrBookNotFound を返す
		}
		return nil, err
	}
	return &book, nil
}

func (r *PostgresRepository) FindAll(ctx context.Context) ([]*domain.Book, error) {
	query := `SELECT id, title, price, isbn FROM books`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Price, &book.ISBN); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
