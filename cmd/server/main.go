package main

import (
	"fmt"
	infrastructure "librigo/internal/infrastructure/repository"
	"librigo/internal/interface/handler"
	"librigo/internal/usecase"
	"net/http"
)

func main() {
	// インフラ層の初期化
	idGen := infrastructure.NewUUIDGenerator()
	repo := infrastructure.NewInMemoryBookRepository()

	// ユースケース層の初期化
	bookUseCase := usecase.NewBookUseCase(repo, idGen)

	// インターフェース層の初期化
	bookHandler := handler.NewBookHandler(bookUseCase)

	// ルーティング設定
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web"))) // 静的ファイルの提供
	mux.HandleFunc("POST /books", bookHandler.Create)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Librigo server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
