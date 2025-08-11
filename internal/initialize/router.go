package initialize

import (
	"net/http"

	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/BaoTo12/go-ecommerce/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.SEVER.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// TODO: Middleware hierarchies as docs/router/router.png
	// r.Use() // Logger
	// r.Use() // Cross
	// r.Use() // Limiter

	// TODO: Mapping router
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/v1/api")
	{
		// Test API
		mainGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Pong Successfully",
			})
		})
		mainGroup.GET("/check-status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Check Status": "Check Status Successfully",
			})
		})
	}
	// TODO: User mapping
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	// TODO: Admin mapping
	{
		adminRouter.InitUserRouter(mainGroup)
		adminRouter.InitAdminRouter(mainGroup)
	}
	return r
}
