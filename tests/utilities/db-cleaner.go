package utilities

import (
	"database/sql"
	"strings"
	"testing"

	viper "github.com/spf13/viper"
)

// DbCleaner ...
func DbCleaner(t *testing.T) {

	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.name")

	connectionDb := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?multiStatements=true"
	db, errConn := sql.Open("mysql", connectionDb)
	if errConn != nil {
		t.Error(errConn)
		return
	}

	getAllTable := "SELECT CONCAT('TRUNCATE TABLE `',table_schema,'`.',TABLE_NAME, ';') AS name FROM INFORMATION_SCHEMA.TABLES WHERE table_schema IN ('" + dbName + "')"
	cleanDB, err := db.Query(getAllTable)
	if err != nil {
		t.Error(err)
		return
	}
	defer cleanDB.Close()

	var tableName string
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	for cleanDB.Next() {
		cleanDB.Scan(&tableName)
		if strings.Contains(tableName, "schema_migrations") {
			continue
		}

		if _, err := db.Exec(tableName); err != nil {
			t.Error(err)
			return
		}
	}
	db.Exec("SET FOREIGN_KEY_CHECKS=1;")
}
