//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/BaoTo12/go-ecommerce/internal/controller"
	"github.com/BaoTo12/go-ecommerce/internal/repo"
	"github.com/BaoTo12/go-ecommerce/internal/service"
	"github.com/google/wire"
)

func InitializeController() *controller.UserController {
	wire.Build(controller.NewUserController, service.NewUserService, repo.NewUserRepository, repo.NewUSerAuthenticationRepository)
	return &controller.UserController{}
}
