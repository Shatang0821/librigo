package auth

import (
	"fmt"
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

func (g *jwtGenerator) Parse(tokenString string) (*user.UserClaims, error) {
	// トークンの解析
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの確認（HS256であることを保証）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return g.secretKey, nil
	})

	if err != nil {
		return nil, user.ErrUnauthorized
	}

	// クレームからユーザー情報を抽出
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, okID := claims["user_id"].(string)
		role, okRole := claims["role"].(string)

		if !okID || !okRole {
			return nil, user.ErrUnauthorized
		}

		return &user.UserClaims{
			UserID: user.UserID(userID),
			Role:   user.UserRole(role),
		}, nil
	}

	return nil, user.ErrUnauthorized
}
