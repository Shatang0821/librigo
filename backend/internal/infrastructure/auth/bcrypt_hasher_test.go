package auth_test

import (
	"librigo/internal/infrastructure/auth"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHasher(t *testing.T) {
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)

	password := "my-secret-password"
	wrongPassword := "not-my-password"

	t.Run("ハッシュ化と照合が成功する", func(t *testing.T) {
		// 1. ハッシュ化
		hash, err := hasher.Hash(password)
		if err != nil {
			t.Fatalf("Hash failed: %v", err)
		}

		if hash == "" || hash == password {
			t.Error("hash should not be empty or equal to the raw password")
		}

		// 2. 正しいパスワードでの照合
		err = hasher.Compare(hash, password)
		if err != nil {
			t.Errorf("Compare should succeed with correct password: %v", err)
		}
	})

	t.Run("間違ったパスワードでの照合が失敗する", func(t *testing.T) {
		hash, _ := hasher.Hash(password)

		// 間違ったパスワードでの照合
		err := hasher.Compare(hash, wrongPassword)
		if err == nil {
			t.Error("Compare should fail with incorrect password")
		}
	})
}
