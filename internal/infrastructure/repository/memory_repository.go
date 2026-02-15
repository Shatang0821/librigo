package repository

import (
	"context"
	"librigo/internal/domain"
	"sync"
)

type InMemoryBookRepository struct {
	mu    sync.RWMutex
	books map[domain.BookID]*domain.Book
}

func NewInMemoryBookRepository() domain.BookRepository {
	return &InMemoryBookRepository{
		books: make(map[domain.BookID]*domain.Book),
	}
}

func (r *InMemoryBookRepository) Save(ctx context.Context, book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) FindByID(ctx context.Context, id domain.BookID) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	book, ok := r.books[id]
	if !ok {
		return nil, nil
	}
	return book, nil
}

func (r *InMemoryBookRepository) FindAll(ctx context.Context) ([]*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []*domain.Book
	for _, b := range r.books {
		list = append(list, b)
	}
	return list, nil
}
