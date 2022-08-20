package dto

import (
	"go-coinstream/pkg/core/entity"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (req *RegisterRequest) ToUserEntity() entity.User {
	return entity.User{
		Email:          req.Email,
		Username:       req.Username,
		HashedPassword: req.Password,
		Name:           req.Name,
	}
}

type RegisterResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func NewRegisterResponse(user *entity.User) *RegisterResponse {
	return &RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}
