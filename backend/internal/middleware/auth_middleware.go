package middleware

import (
	"context"
	"librigo/internal/domain/user"
	"librigo/internal/handler"
	"net/http"
	"strings"
)

type contextKey string

const (
	userClaimsKey contextKey = "user_claims"
)

// AuthMiddleware は JWT を検証するミドルウェアです
func AuthMiddleware(tokenGen user.TokenGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Authorization ヘッダーの取得
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				handler.RespondWithError(w, user.ErrUnauthorized)
				return
			}

			// "Bearer <token>" の形式を確認
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				handler.RespondWithError(w, user.ErrUnauthorized)
				return
			}

			tokenString := parts[1]

			// トークンの解析
			claims, err := tokenGen.Parse(tokenString)
			if err != nil {
				handler.RespondWithError(w, user.ErrUnauthorized)
				return
			}

			// Context に Claims をセットして次の処理へ
			ctx := context.WithValue(r.Context(), userClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetClaims は Context からユーザー情報を取り出すヘルパー関数
func GetClaims(ctx context.Context) (*user.UserClaims, bool) {
	claims, ok := ctx.Value(userClaimsKey).(*user.UserClaims)
	return claims, ok
}
