package api

import (
	"testing"

	viper "github.com/spf13/viper"
)

// InitialTest ...
func InitialTest(t *testing.T) {
	// initial call to envinronment variable
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../../.env.test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err.Error())
	}
}
