package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() // gin.Default() is used to create default instance of gin instance

	return r
}
