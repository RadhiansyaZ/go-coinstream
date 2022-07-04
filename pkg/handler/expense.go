package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/service"
	"io/ioutil"
	"net/http"
)

type ExpenseHandlers interface {
	GetAllExpenses(w http.ResponseWriter, r *http.Request)
	GetExpenseByID(w http.ResponseWriter, r *http.Request)
	CreateExpense(w http.ResponseWriter, r *http.Request)
	UpdateExpense(w http.ResponseWriter, r *http.Request)
	DeleteExpenseByID(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	service service.ExpenseService
}

func NewHttpExpenseHandler(expenseService service.ExpenseService) *Handlers {
	return &Handlers{
		service: expenseService,
	}
}

func (h *Handlers) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	expenses := h.service.FindAll(r.Context())

	result, err := json.Marshal(expenses)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func (h *Handlers) GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	expense, err := h.service.FindById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}
func (h *Handlers) CreateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exp := new(dto.ExpenseRequest)
	_ = json.Unmarshal(body, &exp)

	expense, err := h.service.Add(r.Context(), exp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
func (h *Handlers) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exp := new(dto.ExpenseRequest)
	_ = json.Unmarshal(body, &exp)

	expense, err := h.service.Update(r.Context(), id, exp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}
func (h *Handlers) DeleteExpenseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
