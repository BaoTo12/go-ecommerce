package controller

import (
	"github.com/BaoTo12/go-ecommerce/internal/service"
	"github.com/BaoTo12/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller --> service --> repo --> models --> dbs
func (uc *UserController) GetUserById(ctx *gin.Context) {
	response.SuccessResponse(ctx, 20001, uc.userService.GetUserInformation())
}
