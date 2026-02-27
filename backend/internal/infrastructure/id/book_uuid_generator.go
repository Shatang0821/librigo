package id

import (
	"librigo/internal/domain/book"

	"github.com/google/uuid"
)

type bookUUIDGenerator struct{}

func NewBookUUIDGenerator() book.IDGenerator {
	return &bookUUIDGenerator{}
}

func (g *bookUUIDGenerator) Generate() book.BookID {
	return book.BookID(uuid.New().String())
}
