package initialize

import (
	"fmt"

	c "github.com/BaoTo12/go-ecommerce/internal/controller"
	"github.com/gin-gonic/gin"
)

// ! Plain middleware
func AA(ctx *gin.Context) {
	fmt.Println("Before --- AA")
	ctx.Next()
	fmt.Println("After -- AA")
}
func BB(ctx *gin.Context) {
	fmt.Println("Before --- BB")
	ctx.Next()
	fmt.Println("After -- BB")
}
func CC(ctx *gin.Context) {
	fmt.Println("Before --- CC")
	ctx.Next()
	fmt.Println("After -- CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance

	// TODO: middleware
	r.Use(AA, BB, CC)

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
