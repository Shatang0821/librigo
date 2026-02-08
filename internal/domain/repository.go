package domain

import "context"

type BookRepository interface {
	Svae(ctx context.Context, book *Book) error
	FindByID(ctx context.Context, id BookID) (*Book, error)
	FindAll(ctx context.Context) ([]*Book, error)
}
