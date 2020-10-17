package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/viper"
)

func main() {
	version := flag.String("version", "none", "-")
	state := flag.String("state", "none", "-")
	stage := flag.String("stage", "dev", "-")

	flag.Parse()

	if *version == "none" {
		panic("Stop version must be spellied")
	}

	if *state == "none" {
		panic("Stop state must be spellied")
	}

	// initial call to envinronment variable
	viper.SetConfigType("yaml")

	if *stage == "prod" {
		viper.SetConfigFile(".env.prod.yaml")
	} else if *stage == "test" {
		viper.SetConfigFile(".env.test.yaml")
	} else {
		viper.SetConfigFile(".env.yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbUsername := viper.GetString("databse.username")
	dbPassword := viper.GetString("databse.password")
	dbHost := viper.GetString("databse.host")
	dbPort := viper.GetString("databse.port")
	dbName := viper.GetString("databse.name")

	db, errConn := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?multiStatements=true")
	if errConn != nil {
		fmt.Println(errConn)
		return
	}

	driver, errMysql := mysql.WithInstance(db, &mysql.Config{})
	if errMysql != nil {
		fmt.Println(errMysql)
		return
	}

	m, errInstance := migrate.NewWithDatabaseInstance(
		"file://migrations/"+*version,
		"mysql",
		driver,
	)
	if errInstance != nil {
		fmt.Println(errInstance)
		return
	}

	if *state == "up" {
		if err := m.Up(); err != nil {
			fmt.Println(err, " when up")
		}
	} else if *state == "down" {
		if err := m.Down(); err != nil {
			fmt.Println(err, " when down")
		}
	} else {
		fmt.Println("State is not define right")
	}
}
