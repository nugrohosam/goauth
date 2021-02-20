package databases

import (
	"database/sql"

	"github.com/spf13/viper"
)

// Db using to conn global
var Db *sql.DB

// Conn ..
func Conn() error {
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.name")
	dbDriver := viper.GetString("database.driver")

	dsn := ""
	var db *sql.DB
	var err error

	if dbDriver == "mysql" {
		dsn = dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = sql.Open("mysql", dsn)
	} else if dbDriver == "pgsql" {
		dsn = "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=gorm port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"
		db, err = sql.Open("pgsql", dsn)
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	Db = db

	return nil
}

// Close ..
func Close() error {
	Db.Close()
	return nil
}
