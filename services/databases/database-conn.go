package databases

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB using to conn global
var DB *gorm.DB

// Conn use to connect db gorm
func Conn() error {

	dbUsername := viper.GetString("databse.username")
	dbPassword := viper.GetString("databse.password")
	dbHost := viper.GetString("databse.host")
	dbPort := viper.GetString("databse.port")
	dbName := viper.GetString("databse.name")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
