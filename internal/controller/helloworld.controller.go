package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorldController struct{}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}

func (h *HelloWorldController) HelloWorld(ctx *gin.Context) {
	// Get Param in path
	name := ctx.Param("name")
	// Get Query in path
	// ctx.Query()
	greeting := fmt.Sprintf("Hello %s", name)
	ctx.JSON(http.StatusOK, gin.H{
		"message": greeting,
	})
}
