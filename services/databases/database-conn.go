package databases

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db using to conn global
var Db *gorm.DB

// Conn use to connect db gorm
func Conn() error {
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.name")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, errOpen := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		return errOpen
	}

	sqlDB, errSet := db.DB()
	if errSet != nil {
		return errSet
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	Db = db

	return nil
}
