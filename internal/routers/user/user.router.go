package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// TODO: Non-dependency version

	// ! Public Router
	userPublicRouter := router.Group("/user")
	{
		userPublicRouter.GET("/register")
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
