package initialize

func Run() {
	// load config first
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
