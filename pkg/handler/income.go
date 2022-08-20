package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-coinstream/pkg/core/service"
	"go-coinstream/pkg/handler/dto"
	"io"
	"net/http"
)

type IncomeHandler interface {
	GetAllIncomes(w http.ResponseWriter, r *http.Request)
	GetIncomeByID(w http.ResponseWriter, r *http.Request)
	CreateIncome(w http.ResponseWriter, r *http.Request)
	UpdateIncome(w http.ResponseWriter, r *http.Request)
	DeleteIncomeByID(w http.ResponseWriter, r *http.Request)
}

type IncomeHTTPHandler struct {
	service service.IncomeService
}

func NewHttpIncomeHandler(incomeService service.IncomeService) *IncomeHTTPHandler {
	return &IncomeHTTPHandler{
		service: incomeService,
	}
}

func (i IncomeHTTPHandler) GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	incomes := i.service.FindAll(r.Context())

	result, err := json.Marshal(incomes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (i IncomeHTTPHandler) GetIncomeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	income, err := i.service.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result, err := json.Marshal(income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (i IncomeHTTPHandler) CreateIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inc := new(dto.IncomeRequest)
	_ = json.Unmarshal(body, &inc)

	income, err := i.service.Add(r.Context(), *inc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (i IncomeHTTPHandler) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inc := new(dto.IncomeRequest)
	_ = json.Unmarshal(body, &inc)

	income, err := i.service.Update(r.Context(), id, *inc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(income)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func (i IncomeHTTPHandler) DeleteIncomeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	err := i.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
