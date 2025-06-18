package initialize

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/utils"
	"errors"
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
	global.Mdb = global.Mdb.Debug()
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

// func fromMysqlToGorm() {
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath: "./internal/model",
// 		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
// 	})

// 	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
// 	g.UseDB(global.Mdb) // reuse your gorm db

// 	g.GenerateModel("persons")

// 	//   // Generate basic type-safe DAO API for struct `model.User` following conventions
// 	//   g.ApplyBasic(model.User{})

// 	//   // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
// 	//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

// 	//   // Generate the code
// 	g.Execute()
// }

func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&entity.Account{},
		&entity.Conversation{},
		&entity.Participant{},
		&entity.Profile{},
		&entity.FriendShip{},
		&entity.Message{},
		&entity.MessageRead{},
		&entity.Block{},
		&entity.Role{},
	)
	if err != nil {
		global.Logger.Error(err.Error())
	} else {
		global.Logger.Info("Migrate Table Success")
	}

	initSystemRoles()
	initSystemAdminAccount()
}

func initSystemRoles() {
	var count int64
	global.Mdb.Model(&entity.Role{}).Count(&count)
	if count == 0 {
		roles := []entity.Role{
			{Rolename: "ADMIN"},
			{Rolename: "USER"},
		}
		if err := global.Mdb.Create(&roles).Error; err != nil {
			return
		}
		global.Logger.Info("Seeded roles: ADMIN, USER")
	} else {
		global.Logger.Info("Roles already seeded")
	}
}

func initSystemAdminAccount() {
	var g = global.Config.Admin
	var adminAccount entity.Account
	res := global.Mdb.Model(&entity.Account{}).Where("username = ?", "admin").First(&adminAccount)
	if res.Error == nil {
		global.Logger.Info("Admin account already exists")
		return
	}
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		hashedPassword, err := utils.HashPassword(g.Password)
		if err != nil {
			global.Logger.Error("Failed to hash password for admin")
			return
		}

		var adminRole entity.Role
		if err := global.Mdb.Where("rolename = ?", "ADMIN").First(&adminRole).Error; err != nil {
			global.Logger.Error("Admin role not found")
			return
		}

		account := entity.Account{
			Email:       g.Email,
			Username:    g.Username,
			PhoneNumber: g.PhoneNumber,
			Password:    hashedPassword,
			Roles:       []entity.Role{adminRole}, // phải gán role thật, không tạo mới
			Profile: &entity.Profile{
				FullName:  "Admin",
				Bio:       "",
				AvatarURL: "",
				CoverURL:  "",
			},
		}

		if err := global.Mdb.Create(&account).Error; err != nil {
			global.Logger.Error("Failed to create admin account: " + err.Error())
			return
		}

		global.Logger.Info("✅ Seeded admin account: admin")
	} else {
		global.Logger.Error("Error when checking admin account: " + res.Error.Error())
	}

}
