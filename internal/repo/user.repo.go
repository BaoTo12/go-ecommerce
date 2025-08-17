package repo

import (
	"context"

	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/BaoTo12/go-ecommerce/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) bool {
	result, err := ur.sqlc.GetUserByEmail(ctx, email)
	if err != nil {
		return false
	}
	return result.UsrID != 0
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
