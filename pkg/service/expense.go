package service

import (
	"context"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/entity"
	"go-coinstream/pkg/repository"
)

type ExpenseService interface {
	Add(ctx context.Context, expense *dto.ExpenseRequest) (*entity.Expense, error)
	FindAll(ctx context.Context) []entity.Expense
	FindById(ctx context.Context, id string) (*entity.Expense, error)
	Update(ctx context.Context, id string, expense *dto.ExpenseRequest) (*entity.Expense, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(ctx context.Context, data *dto.ExpenseRequest) (*entity.Expense, error) {
	expense := data.ToExpenseEntity()

	res, err := s.repo.Add(ctx, expense)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) FindAll(ctx context.Context) []entity.Expense {
	res, _ := s.repo.FindAll(ctx)
	return res
}

func (s *Service) FindById(ctx context.Context, id string) (*entity.Expense, error) {
	res, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) Update(ctx context.Context, id string, data *dto.ExpenseRequest) (*entity.Expense, error) {
	_, err := s.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	expense := data.ToExpenseEntity()

	res, err := s.repo.Update(ctx, id, expense)
	return res, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	_, err := s.FindById(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
