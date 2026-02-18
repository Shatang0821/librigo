package repository

import (
	"librigo/internal/domain/book"

	"github.com/google/uuid"
)

type uuidGenerator struct{}

func NewUUIDGenerator() book.IDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) Generate() book.BookID {
	return book.BookID(uuid.New().String())
}
