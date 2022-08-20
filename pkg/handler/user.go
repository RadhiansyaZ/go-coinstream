package handler

import (
	"encoding/json"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/service"
	"io"
	"net/http"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

func NewHttpUserHandler(service service.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dtoRegister dto.RegisterRequest
	_ = json.Unmarshal(body, &dtoRegister)

	user, err := h.service.Register(r.Context(), dtoRegister)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	result, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dtoLogin dto.LoginRequest
	_ = json.Unmarshal(body, &dtoLogin)

	token, err := h.service.Login(r.Context(), dtoLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	result, err := json.Marshal(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
