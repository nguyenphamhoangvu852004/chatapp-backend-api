package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	// fromMysqlToGorm()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	checkErr(err, "Set Mysql Pool Failed")

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func fromMysqlToGorm() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	g.GenerateModel("persons")

	//   // Generate basic type-safe DAO API for struct `model.User` following conventions
	//   g.ApplyBasic(model.User{})

	//   // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	//   // Generate the code
	g.Execute()
}

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		// &po.Role{},
		// &po.User{},
		&model.Person2{},
	)
	if err != nil {
		global.Logger.Error(err.Error())
	} else {
		global.Logger.Info("Migrate Table Success")
	}

}
