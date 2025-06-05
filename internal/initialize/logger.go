package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
