package handler

import (
	"encoding/json"
	"errors"
	"librigo/internal/domain"
	"librigo/internal/usecase"
	"net/http"
)

type BookHandler struct {
	useCase usecase.BookUseCase
}

func NewBookHandler(uc usecase.BookUseCase) *BookHandler {
	return &BookHandler{useCase: uc}
}

// CreateBookRequest は書籍登録のリクエストデータです
type CreateBookRequest struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	ISBN  string `json:"isbn"`
}

// GetBookResponse は書籍登録時のレスポンスデータです
type CreateBookResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// BookResponse は書籍のレスポンスデータです
type BookResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	ISBN  string `json:"isbn"`
}

// Create は書籍登録のHTTPハンドラーです
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	input := usecase.CreateBookInput{
		Title: req.Title,
		Price: req.Price,
		ISBN:  req.ISBN,
	}

	output, err := h.useCase.CreateBook(r.Context(), input)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidBookTitle) || errors.Is(err, domain.ErrInvalidBookPrice) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res := CreateBookResponse{
		ID:    output.ID,
		Title: output.Title,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

// List は書籍一覧取得のHTTPハンドラーです
func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	outputs, err := h.useCase.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := make([]BookResponse, len(outputs))
	for i, o := range outputs {
		res[i] = toBookResponse(o)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	output, err := h.useCase.GetBookByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := toBookResponse(output)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// toBookResponse は usecase.BookOutput を BookResponse に変換するヘルパー関数です
func toBookResponse(book *usecase.BookOutput) BookResponse {
	return BookResponse{
		ID:    book.ID,
		Title: book.Title,
		Price: book.Price,
		ISBN:  book.ISBN,
	}
}
