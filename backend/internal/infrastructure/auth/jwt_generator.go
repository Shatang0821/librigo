package auth

import (
	"librigo/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtGenerator struct {
	secretKey []byte
}

func NewJWTGenerator(secretKey string) user.TokenGenerator {
	return &jwtGenerator{secretKey: []byte(secretKey)}
}

func (g *jwtGenerator) Generate(u *user.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"role":    u.Role,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(g.secretKey)
}
