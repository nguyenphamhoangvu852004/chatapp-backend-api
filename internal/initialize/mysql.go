package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/po"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErr(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg)
		panic(err)
	}
}
func InitMysql() {
	global.Logger.Info("Init mysql...")
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErr(err, "Init mysql failed")
	global.Mdb = db
	global.Logger.Info("MysqlPool Initialize Successfully")

	// setPool
	SetPool()
	migrateTable()

}

func SetPool() {

	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(m.ConnMaxLifeTime))
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	checkErr(err, "Fail to migrate table")
}
