package databases

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"

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
	} else if dbDriver == "postgres" {
		dsn = "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?" + "sslmode=disableTimeZone=Asia/Jakarta"
		db, err = sql.Open("postgres", dsn)
	} else {
		return errors.New("Not Defined Database Driver")
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// POOLING DB
	dbMinPool := viper.GetInt("database.min_pool")
	dbMaxPool := viper.GetInt("database.max_pool")

	db.SetMaxIdleConns(dbMinPool)
	db.SetMaxOpenConns(dbMaxPool)
	db.SetConnMaxLifetime(time.Hour)

	Db = db

	return nil
}

// Close ..
func Close() error {
	Db.Close()
	return nil
}
