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

type CreateBookRequest struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	ISBN  string `json:"isbn"`
}

type CreateBookResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

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
