package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) Ping(ctx *gin.Context) {
	fmt.Println("My Pong ----->")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping pong..",
	})
}
