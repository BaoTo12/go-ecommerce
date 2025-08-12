package service

import (
	"github.com/BaoTo12/go-ecommerce/internal/repo"
	"github.com/BaoTo12/go-ecommerce/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// check whether email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrUserAlreadyExisted
	}
	return response.ErrCodeSuccess
}
