package repository

import (
	"context"
	"database/sql"
	"go-coinstream/pkg/entity"
)

type IncomeRepository interface {
	Add(ctx context.Context, income *entity.Income) (*entity.Income, error)
	FindAll(ctx context.Context) ([]entity.Income, error)
	FindById(ctx context.Context, id string) (*entity.Income, error)
	Update(ctx context.Context, id string, income *entity.Income) (*entity.Income, error)
	Delete(ctx context.Context, id string) error
}

type IncomeStore struct {
	db *sql.DB
}

func NewIncomeRepository(db *sql.DB) *IncomeStore {
	return &IncomeStore{db: db}
}

func (s *IncomeStore) Add(ctx context.Context, income *entity.Income) (*entity.Income, error) {
	sqlStatement := `INSERT INTO income(name,amount,date) VALUES($1,$2,$3) RETURNING id`

	row := s.db.QueryRowContext(ctx, sqlStatement, income.Name, income.Amount, income.Date)

	var insertId string
	err := row.Scan(&insertId)
	if err != nil {
		return nil, err
	}

	income.ID = insertId

	return income, nil
}

func (s *IncomeStore) FindAll(ctx context.Context) ([]entity.Income, error) {
	sqlStatement := `SELECT id, name, amount, date FROM income`

	rows, err := s.db.QueryContext(ctx, sqlStatement)
	if err == sql.ErrNoRows {
		return nil, err
	}

	var incomes []entity.Income

	for rows.Next() {
		var inc entity.Income
		if err := rows.Scan(&inc.ID, &inc.Name, &inc.Amount, &inc.Date); err != nil {
			return incomes, err
		}
		incomes = append(incomes, inc)
	}

	if err = rows.Err(); err != nil {
		return incomes, err
	}

	return incomes, nil
}

func (s *IncomeStore) FindById(ctx context.Context, id string) (*entity.Income, error) {
	sqlStatement := `SELECT id, name, amount, date FROM income where id=$1`

	var income entity.Income

	row := s.db.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&income.ID, &income.Name, &income.Amount, &income.Date)
	if err != nil {
		return nil, err
	}

	return &income, nil
}

func (s *IncomeStore) Update(ctx context.Context, id string, income *entity.Income) (*entity.Income, error) {
	sqlStatement := `UPDATE income 
						SET NAME=$2,
							AMOUNT=$3,
							DATE=$4
						WHERE id=$1
						RETURNING *`

	row := s.db.QueryRowContext(ctx, sqlStatement, id, income.Name, income.Amount, income.Date)

	err := row.Scan(&income.ID, &income.Name, &income.Amount, &income.Date)
	if err != nil {
		return nil, err
	}

	return income, nil
}

func (s *IncomeStore) Delete(ctx context.Context, id string) error {
	sqlStatement := `DELETE FROM income WHERE id=$1`

	s.db.QueryRowContext(ctx, sqlStatement, id)
	return nil
}
