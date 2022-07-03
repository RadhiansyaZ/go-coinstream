package dto

import "go-coinstream/pkg/entity"

type ExpenseRequest struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

func (expense *ExpenseRequest) ToExpenseEntity() *entity.Expense {
	return &entity.Expense{
		Name:     expense.Name,
		Amount:   expense.Amount,
		Category: expense.Category,
		Date:     expense.Date,
	}
}

type ExpenseResponse struct {
	ID       string
	Name     string
	Amount   float64
	Category string
	Date     string
}
