package routers

import (
	c "github.com/BaoTo12/go-ecommerce/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", c.NewPongController().Ping)
		v1.GET("/hello/:name", c.NewHelloWorldController().HelloWorld)
		v1.GET("/user", c.NewUserController().GetUserById)
	}
	return r
}
