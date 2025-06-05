package global

import (
	"chapapp-backend-api/pkg/logger"
	"chapapp-backend-api/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb   *gorm.DB
)

/*
redis
mysql
...
*/
