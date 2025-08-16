package controller

import (
	"fmt"

	"github.com/BaoTo12/go-ecommerce/internal/service"
	"github.com/BaoTo12/go-ecommerce/internal/vo"
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
func (uc *UserController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var bodyRequest vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		fmt.Println(err)
		response.FailureResponse(c, response.ErrCodeParamInvalid)
		return
	}
	result := uc.userService.Register(ctx, bodyRequest.Email, bodyRequest.Purpose)
	response.SuccessResponse(c, result, nil)
}
