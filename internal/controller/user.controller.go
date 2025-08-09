package controller

import (
	"net/http"

	"github.com/BaoTo12/go-ecommerce/internal/service"
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
	ctx.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUserInformation(),
	})
}
