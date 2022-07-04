package service

import (
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/entity"
	"go-coinstream/pkg/repository"
)

type ExpenseService interface {
	Add(expense *dto.ExpenseRequest) (*entity.Expense, error)
	FindAll() []entity.Expense
	FindById(id string) (*entity.Expense, error)
	Update(id string, expense *dto.ExpenseRequest) (*entity.Expense, error)
	Delete(id string) error
}

type Service struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(data *dto.ExpenseRequest) (*entity.Expense, error) {
	expense := data.ToExpenseEntity()

	res, err := s.repo.Add(expense)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) FindAll() []entity.Expense {
	res, _ := s.repo.FindAll()
	return res
}

func (s *Service) FindById(id string) (*entity.Expense, error) {
	res, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) Update(id string, data *dto.ExpenseRequest) (*entity.Expense, error) {
	_, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	expense := data.ToExpenseEntity()

	res, err := s.repo.Update(id, expense)
	return res, nil
}

func (s *Service) Delete(id string) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
