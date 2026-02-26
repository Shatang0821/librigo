package handler

import (
	"encoding/json"
	"librigo/internal/domain/apperror"
	"librigo/internal/usecase"
	"net/http"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: uc}
}

// SignUpRequest はユーザー登録のリクエストデータです
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 一時的に
		RespondWithError(w, apperror.ErrInvalidJSON)
		return
	}

	input := usecase.SignUpInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.useCase.SignUp(r.Context(), input)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	res := SignUpResponse{
		ID:    output.ID,
		Name:  output.Name,
		Email: output.Email,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
