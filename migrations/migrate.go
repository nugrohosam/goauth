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

	flag.Parse()

	if *version == "none" {
		panic("Stop version must be spellied")
	}

	if *state == "none" {
		panic("Stop state must be spellied")
	}

	dbUsername := viper.GetString("databse.username")
	dbPassword := viper.GetString("databse.password")
	dbHost := viper.GetString("databse.host")
	dbPort := viper.GetString("databse.port")
	dbName := viper.GetString("databse.name")

	db, _ := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+*version,
		"mysql",
		driver,
	)

	if *state == "up" {
		m.Up()
	} else if *state == "down" {
		m.Down()
	} else {
		fmt.Println("State is not define right")
	}
}
