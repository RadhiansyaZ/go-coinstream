package service

import (
	"context"
	"errors"
	"go-coinstream/pkg/dto"
	"go-coinstream/pkg/repository"
	"go-coinstream/pkg/service/util"
)

type UserService interface {
	Register(ctx context.Context, data dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, data dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) Register(ctx context.Context, data dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user := data.ToUserEntity()

	hashedPwd, err := util.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user.HashedPassword = hashedPwd

	res, err := s.repository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	response := dto.NewRegisterResponse(res)

	return response, nil
}

func (s *userService) Login(ctx context.Context, data dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repository.FindByUsername(ctx, data.Username)
	if err != nil {
		return nil, errors.New("username and password didn't match")
	}

	if !util.CheckPasswordHash(data.Password, user.HashedPassword) {
		return nil, errors.New("username and password didn't match")
	}

	token, err := util.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	response := dto.LoginResponse{AccessToken: token}
	return &response, nil
}
