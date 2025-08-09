package service

import "github.com/BaoTo12/go-ecommerce/internal/repo"

type UserService struct {
	UserRepository *repo.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repo.NewUserRepository(),
	}
}

func (us *UserService) GetUserInformation() string {
	return us.UserRepository.GetUserInformation()
}
