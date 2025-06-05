package initialize

import (
	"chapapp-backend-api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Load Config Success")
	InitMysql()
	InitRedis()
	InitRouter().Run(":8080")
}
