package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ping pong..",
	})
}
