package dto

import "go-coinstream/pkg/entity"

type IncomeRequest struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

func (income *IncomeRequest) ToIncomeEntity() *entity.Income {
	return &entity.Income{
		Name:   income.Name,
		Amount: income.Amount,
		Date:   income.Date,
	}
}
