package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

// CanAccessBy using for ..
func CanAccessBy(s []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := helpers.GetSessionDataString(c.Request, c.Writer, viper.GetString("config.session.userID"))
		userIDInt, _ := strconv.Atoi(userID)
		if isExists, err := usecases.IsHaveAnyRole(userIDInt, s); !isExists || err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr(err.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}
