package main

import (
	database "github.com/nugrohosam/gosampleapi/services/databases"
	httpConn "github.com/nugrohosam/gosampleapi/services/http"
	viper "github.com/spf13/viper"
)

func main() {

	// initial call to envinronment variable
	viper.SetConfigType("yaml")
	viper.SetConfigFile(".env.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := database.Conn(); err != nil {
		panic(err)
	}

	if err := httpConn.Serve(); err != nil {
		panic(err)
	}
}
