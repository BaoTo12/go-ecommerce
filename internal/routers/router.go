package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance
	/// gin.Context
	/*
		gin.Context is a Gin's struct which is designed to encapsulate everything you need during the lifecycle of an HTTP request
	*/
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", Ping)
		v1.GET("/hello/:name", HelloWorld)
	}
	return r
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping pong..",
	})
}

func HelloWorld(ctx *gin.Context) {
	// Get Param in path
	name := ctx.Param("name")
	// Get Query in path
	// ctx.Query()
	greeting := fmt.Sprintf("Hello %s", name)
	ctx.JSON(http.StatusOK, gin.H{
		"message": greeting,
	})
}
