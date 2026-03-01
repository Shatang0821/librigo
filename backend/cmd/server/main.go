package main

import (
	"fmt"
	"librigo/internal/handler"
	"librigo/internal/infrastructure/auth"
	"librigo/internal/infrastructure/database"
	"librigo/internal/infrastructure/id"
	"librigo/internal/infrastructure/postgres"
	"os"

	"librigo/internal/usecase"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CORSミドルウェアの定義
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// フロントエンドのURLを許可
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// プリフライトリクエスト（OPTIONS）への対応
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// データベース接続の初期化
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// インフラ層の初期化
	bookIdGen := id.NewBookUUIDGenerator()
	userIdGen := id.NewUserUUIDGenerator()
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-very-secret-key" // 開発用デフォルト
	}
	tokenGen := auth.NewJWTGenerator(jwtSecret)
	// リポジトリ層の初期化
	bookRepo := postgres.NewBookRepository(db)
	userRepo := postgres.NewUserRepository(db)
	// ユースケース層の初期化
	bookUseCase := usecase.NewBookUseCase(bookRepo, bookIdGen)
	userUseCase := usecase.NewUserUseCase(userRepo, hasher, userIdGen, tokenGen)
	// ハンドラ層の初期化
	bookHandler := handler.NewBookHandler(bookUseCase)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティング設定
	mux := http.NewServeMux()

	// 書籍登録
	mux.HandleFunc("POST /books", bookHandler.Create)
	// 書籍一覧取得
	mux.HandleFunc("GET /books", bookHandler.List)
	// 書籍詳細取得
	mux.HandleFunc("GET /books/{id}", bookHandler.Get)

	// ユーザー登録
	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("POST /signin", userHandler.SignIn)
	// ヘルスチェック
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Librigo server starting on :8080...")
	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		panic(err)
	}
}
