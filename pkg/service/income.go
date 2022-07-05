package service

import (
	"context"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/entity"
	"go-coinstream/pkg/repository"
)

type IncomeService interface {
	Add(ctx context.Context, data *dto.IncomeRequest) (*entity.Income, error)
	FindAll(ctx context.Context) []entity.Income
	FindById(ctx context.Context, id string) (*entity.Income, error)
	Update(ctx context.Context, id string, data *dto.IncomeRequest) (*entity.Income, error)
	Delete(ctx context.Context, id string) error
}

type incomeService struct {
	repo repository.IncomeRepository
}

func NewIncomeService(repo repository.IncomeRepository) *incomeService {
	return &incomeService{repo: repo}
}

func (s *incomeService) Add(ctx context.Context, data *dto.IncomeRequest) (*entity.Income, error) {
	income := data.ToIncomeEntity()

	res, err := s.repo.Add(ctx, income)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *incomeService) FindAll(ctx context.Context) []entity.Income {
	res, _ := s.repo.FindAll(ctx)

	return res
}

func (s *incomeService) FindById(ctx context.Context, id string) (*entity.Income, error) {
	res, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *incomeService) Update(ctx context.Context, id string, data *dto.IncomeRequest) (*entity.Income, error) {
	_, err := s.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	income := data.ToIncomeEntity()

	res, err := s.repo.Update(ctx, id, income)
	return res, nil
}

func (s *incomeService) Delete(ctx context.Context, id string) error {
	_, err := s.FindById(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
