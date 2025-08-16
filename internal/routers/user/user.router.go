package user

import (
	"github.com/BaoTo12/go-ecommerce/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// TODO: Initialize controller
	userController := wire.InitializeController()
	// ! Public Router
	userPublicRouter := router.Group("/user")
	{
		userPublicRouter.POST("/register", userController.Register)
		userPublicRouter.POST("/otp")
	}
	// ! Private Router

	userPrivateRouter := router.Group("/user")
	// userPrivateRouter.Use(Limiter())
	// userPrivateRouter.Use(Auth())
	// userPrivateRouter.Use(Permission())
	{
		userPrivateRouter.GET("/user-info")
	}
}
