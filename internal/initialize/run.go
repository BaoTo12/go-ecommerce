package initialize

import (
	"fmt"

	"github.com/BaoTo12/go-ecommerce/global"
)

func Run() {
	// load config first
	LoadConfig()
	fmt.Println("Loading configuration ", global.Config.MYSQL)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run("127.0.0.1:8002")
}
