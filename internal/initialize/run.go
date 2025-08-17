package initialize

import (
	"fmt"

	"github.com/BaoTo12/go-ecommerce/global"
)

func Run() {
	// load config first
	LoadConfig()
	fmt.Println("Loading configuration ", global.Config.LOGGER)
	InitLogger()
	global.Logger.Info("Logger is ok!!")
	// InitMysql()
	InitMysqlCompile()
	InitRedis()
	r := InitRouter()

	serverSetting := global.Config.SEVER
	r.Run(fmt.Sprintf("%s:%v", serverSetting.Host, serverSetting.Port))
}
