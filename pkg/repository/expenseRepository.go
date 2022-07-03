package repository

import (
	"database/sql"
	"go-coinstream/pkg/entity"
	"log"
)

type ExpenseRepository interface {
	Add(expense *entity.Expense) (*entity.Expense, error)
	FindAll() ([]entity.Expense, error)
	FindById(id string) (*entity.Expense, error)
	Update(id string, expense *entity.Expense) (*entity.Expense, error)
	Delete(id string) error
}

type ExpenseStore struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseStore {
	return &ExpenseStore{db: db}
}
func (s *ExpenseStore) Add(expense *entity.Expense) (*entity.Expense, error) {
	sqlStatement := `INSERT INTO expense(name,amount,category,date) VALUES($1,$2,$3,$4) RETURNING id`

	row := s.db.QueryRow(sqlStatement, expense.Name, expense.Amount, expense.Category, expense.Date)

	var insertId string
	err := row.Scan(&insertId)
	if err != nil {
		return nil, err
	}

	expense.ID = insertId

	return expense, nil
}

func (s *ExpenseStore) FindAll() ([]entity.Expense, error) {
	sqlStatement := `SELECT * FROM expense`

	rows, err := s.db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	var expenses []entity.Expense

	for rows.Next() {
		var exp entity.Expense
		if err := rows.Scan(&exp.ID, &exp.Name, &exp.Amount, &exp.Category, &exp.Date); err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}

	if err = rows.Err(); err != nil {
		return expenses, err
	}

	return expenses, nil
}
func (s *ExpenseStore) FindById(id string) (*entity.Expense, error) {
	sqlStatement := `SELECT * FROM expense WHERE id=$1`

	var expense entity.Expense

	row := s.db.QueryRow(sqlStatement, id)

	err := row.Scan(&expense.ID, &expense.Name, &expense.Amount, &expense.Category, &expense.Date)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}
func (s *ExpenseStore) Update(id string, expense *entity.Expense) (*entity.Expense, error) {
	sqlStatement := `UPDATE expense
					 SET NAME=$2,
					     AMOUNT=$3,
					     CATEGORY=$4,
					     DATE=$5
					 WHERE id=$1
					 RETURNING *`

	row := s.db.QueryRow(sqlStatement, id, expense.Name, expense.Amount, expense.Category, expense.Date)

	err := row.Scan(&expense.ID, &expense.Name, &expense.Amount, &expense.Category, &expense.Date)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *ExpenseStore) Delete(id string) error {
	sqlStatement := `DELETE FROM expense WHERE id=$1`

	s.db.QueryRow(sqlStatement, id)
	return nil
}
