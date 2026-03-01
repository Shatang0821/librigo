package auth

import (
	"librigo/internal/domain/user"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWTGenerator_Generate(t *testing.T) {
	secret := "test-secret-key"
	gen := NewJWTGenerator(secret)

	// テストデータ
	testUser := &user.User{
		ID:   "user-123",
		Role: user.RoleMember,
	}

	duration := time.Hour

	tokenString, err := gen.Generate(testUser, duration)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// UserIDの検証
		if claims["user_id"] != string(testUser.ID) {
			t.Errorf("got user_id %v, want %v", claims["user_id"], testUser.ID)
		}
		// Roleの検証
		if claims["role"] != string(testUser.Role) {
			t.Errorf("got role %v, want %v", claims["role"], testUser.Role)
		}
		// 有効期限の検証（だいたい1時間後になっているか）
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			t.Error("token is already expired")
		}
	} else {
		t.Error("invalid claims or token")
	}
}
