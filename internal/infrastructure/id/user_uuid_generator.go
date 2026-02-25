package id

import (
	"librigo/internal/domain/user"

	"github.com/google/uuid"
)

type userUUIDGenerator struct{}

func NewUserUUIDGenerator() user.IDGenerator {
	return &userUUIDGenerator{}
}

func (g *userUUIDGenerator) Generate() user.UserID {
	return user.UserID(uuid.New().String())
}
