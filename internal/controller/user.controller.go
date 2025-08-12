package controller

import (
	"github.com/BaoTo12/go-ecommerce/internal/service"
	"github.com/BaoTo12/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// controller --> service --> repo --> models --> dbs
func (uc *UserController) Register(ctx *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(ctx, result, nil)
}
