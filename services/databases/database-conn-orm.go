package databases

import (
	"errors"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbOrm using to conn global
var DbOrm *gorm.DB

// ConnOrm ..
func ConnOrm() error {
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.name")
	dbDriver := viper.GetString("database.driver")

	dsn := ""
	var db *gorm.DB
	var errOpen error

	if dbDriver == "mysql" {
		dsn = dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, errOpen = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbDriver == "postgres" {
		dsn = "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?" + "TimeZone=Asia/Jakarta"
		db, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		return errors.New("Not Defined Database Driver")
	}

	if errOpen != nil {
		return errOpen
	}

	sqlDB, errSet := db.DB()
	if errSet != nil {
		return errSet
	}

	// POOLING DB
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DbOrm = db

	return nil
}
