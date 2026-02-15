package repository

import (
	"librigo/internal/domain"

	"github.com/google/uuid"
)

type uuidGenerator struct{}

func NewUUIDGenerator() domain.IDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) Generate() domain.BookID {
	return domain.BookID(uuid.New().String())
}
