package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// ! Public Router
	adminPublicRouter := router.Group("/admin")
	{
		adminPublicRouter.POST("/login")
	}

	// ! Private Router
	// adminPrivateRouter := router.Group("/admin")
	// userPrivateRouter.Use(Limiter())
	// userPrivateRouter.Use(Auth())
	// userPrivateRouter.Use(Permission())
}
