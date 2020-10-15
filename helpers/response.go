package helpers

import (
	"github.com/gin-gonic/gin"
	viper "github.com/spf13/viper"
)

// ResponseErrValidation ...
func ResponseErrValidation(errors []map[string]interface{}) gin.H {
	resp := gin.H{
		"version": viper.GetString("app.version"),
		"message": nil,
		"errors":  errors,
	}

	return resp
}

// ResponseErr ...
func ResponseErr(message string) gin.H {
	resp := gin.H{
		"version": viper.GetString("app.version"),
		"message": message,
		"errors":  nil,
	}

	return resp
}

// ResponseMany ...
func ResponseMany(data []map[string]interface{}) gin.H {
	resp := gin.H{
		"version": viper.GetString("app.version"),
		"items":   data,
	}

	return resp
}

// ResponseManyWithPagination ...
func ResponseManyWithPagination(data []map[string]interface{}, pagination []map[string]interface{}) gin.H {
	resp := gin.H{
		"version":    viper.GetString("app.version"),
		"data":       data,
		"pagination": pagination,
	}

	return resp
}

// ResponseOne ...
func ResponseOne(data []map[string]interface{}) gin.H {
	resp := gin.H{
		"version": viper.GetString("app.version"),
		"data":    data,
	}

	return resp
}
