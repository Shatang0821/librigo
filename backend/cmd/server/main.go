package main

import (
	"fmt"
	"librigo/internal/handler"
	"librigo/internal/infrastructure/auth"
	"librigo/internal/infrastructure/database"
	"librigo/internal/infrastructure/id"
	"librigo/internal/infrastructure/postgres"

	"librigo/internal/usecase"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
	// リポジトリ層の初期化
	bookRepo := postgres.NewBookRepository(db)
	userRepo := postgres.NewUserRepository(db)
	// ユースケース層の初期化
	bookUseCase := usecase.NewBookUseCase(bookRepo, bookIdGen)
	userUseCase := usecase.NewUserUseCase(userRepo, hasher, userIdGen)
	// ハンドラ層の初期化
	bookHandler := handler.NewBookHandler(bookUseCase)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティング設定
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web"))) // 静的ファイルの提供
	// 書籍登録
	mux.HandleFunc("POST /books", bookHandler.Create)
	// 書籍一覧取得
	mux.HandleFunc("GET /books", bookHandler.List)
	// 書籍詳細取得
	mux.HandleFunc("GET /books/{id}", bookHandler.Get)

	// ユーザー登録
	mux.HandleFunc("POST /signup", userHandler.SignUp)

	// ヘルスチェック
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Librigo server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
