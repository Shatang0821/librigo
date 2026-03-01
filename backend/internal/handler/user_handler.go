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
type signUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string          `json:"token"`
	User  *signUpResponse `json:"user"`
}

// ユーザー登録
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpRequest
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

	res := signUpResponse{
		ID:    output.ID,
		Name:  output.Name,
		Email: output.Email,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, apperror.ErrInvalidJSON)
		return
	}

	input := usecase.SignInInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.useCase.SignIn(r.Context(), input)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	res := signInResponse{
		Token: output.Token,
		User: &signUpResponse{
			ID:    output.User.ID,
			Name:  output.User.Name,
			Email: output.User.Email,
		},
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
