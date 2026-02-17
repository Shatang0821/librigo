package main

import (
	"fmt"
	"librigo/internal/infrastructure/database"
	"librigo/internal/infrastructure/repository"
	"librigo/internal/interface/handler"
	"librigo/internal/usecase"
	"log"
	"net/http"
)

func main() {
	// データベース接続の初期化
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	// インフラ層の初期化
	idGen := repository.NewUUIDGenerator()
	repo := repository.NewPostgresRepository(db)

	// ユースケース層の初期化
	bookUseCase := usecase.NewBookUseCase(repo, idGen)

	// インターフェース層の初期化
	bookHandler := handler.NewBookHandler(bookUseCase)

	// ルーティング設定
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web"))) // 静的ファイルの提供
	// 書籍登録
	mux.HandleFunc("POST /books", bookHandler.Create)
	// 書籍一覧取得
	mux.HandleFunc("GET /books", bookHandler.List)
	// 書籍詳細取得
	mux.HandleFunc("GET /books/{id}", bookHandler.Get)

	// ヘルスチェック
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Librigo server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
