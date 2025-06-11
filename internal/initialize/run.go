package initialize

import (
	"chapapp-backend-api/global"
	"strconv"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Load Config Success")
	InitMysql()
	InitRouter().Run(":" + strconv.Itoa(global.Config.Server.Port))
}
