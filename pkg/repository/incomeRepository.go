package repository

import (
	"database/sql"
	"go-coinstream/pkg/entity"
)

type IncomeRepository interface {
	Add(expense *entity.Income) (*entity.Income, error)
	FindAll() []entity.Income
	FindById(id string) (*entity.Income, error)
	Update(id string, expense *entity.Income) (*entity.Income, error)
	Delete(id string) error
}

type IncomeStore struct {
	db *sql.DB
}
