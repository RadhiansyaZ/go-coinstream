package repository

import (
	"context"
	"database/sql"
	"go-coinstream/pkg/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user entity.User) (*entity.User, error) {
	sqlStatement := `INSERT INTO user_coinstream(email,username,hashed_password,name) VALUES($1,$2,$3,$4) RETURNING id, created_at`

	var userId string
	var createdAt string

	err := r.db.QueryRowContext(ctx, sqlStatement, (&user).Email, (&user).Username, (&user).HashedPassword, (&user).Name).Scan(&userId, &createdAt)

	if err != nil {
		return nil, err
	}

	user.ID = userId
	user.CreatedAt = createdAt

	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	sqlStatement := `SELECT * FROM user_coinstream WHERE username=$1`

	var user entity.User

	err := r.db.QueryRowContext(ctx, sqlStatement, username).Scan(&user.ID, &user.Email, &user.Username, &user.HashedPassword, &user.Name, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
