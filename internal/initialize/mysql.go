package initialize

import (
	"fmt"
	"time"

	"github.com/longtk26/go-ecommerce/global"
	"github.com/longtk26/go-ecommerce/internal/po"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func checkErrorPanic(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)

  	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,	// skip transaction to improve performance
	})

	checkErrorPanic(err, "Failed to init MYSQL")
	global.Logger.Info("MYSQL connected")
	global.Mdb = db

	// Set pool
	SetPool()

	// Migrate tables
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()

	if err != nil {
		fmt.Printf("Failed to set pool: %v", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	checkErrorPanic(err, "Failed to migrate tables")
}