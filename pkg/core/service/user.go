package service

import (
	"context"
	"errors"
	util2 "go-coinstream/pkg/core/service/util"
	dto2 "go-coinstream/pkg/handler/dto"
	"go-coinstream/pkg/repository"
)

type UserService interface {
	Register(ctx context.Context, data dto2.RegisterRequest) (*dto2.RegisterResponse, error)
	Login(ctx context.Context, data dto2.LoginRequest) (*dto2.LoginResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) Register(ctx context.Context, data dto2.RegisterRequest) (*dto2.RegisterResponse, error) {
	user := data.ToUserEntity()

	hashedPwd, err := util2.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user.HashedPassword = hashedPwd

	res, err := s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	response := dto2.NewRegisterResponse(res)

	return response, nil
}

func (s *userService) Login(ctx context.Context, data dto2.LoginRequest) (*dto2.LoginResponse, error) {
	user, err := s.repository.FindByUsername(ctx, data.Username)
	if err != nil {
		return nil, errors.New("username and password didn't match")
	}

	if !util2.CheckPasswordHash(data.Password, user.HashedPassword) {
		return nil, errors.New("username and password didn't match")
	}

	token, err := util2.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	response := dto2.LoginResponse{AccessToken: token}
	return &response, nil
}
