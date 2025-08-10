package routers

import (
	c "github.com/BaoTo12/go-ecommerce/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance

	pongController := c.NewPongController()
	HelloWorldController := c.NewHelloWorldController()
	userController := c.NewUserController()

	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", pongController.Ping)
		v1.GET("/hello/:name", HelloWorldController.HelloWorld)
		v1.GET("/user", userController.GetUserById)
	}
	return r
}
