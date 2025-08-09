package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance
	/// gin.Context
	/*
		gin.Context is a Gin's struct which is designed to encapsulate everything you need during the lifecycle of an HTTP request
	*/
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", Ping)
	}
	r.Run()
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping pong..",
	})
}
