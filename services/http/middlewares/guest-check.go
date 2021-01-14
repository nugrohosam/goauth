package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	helpers "github.com/nugrohosam/gosampleapi/helpers"
)

// GuestMiddlewre ..
var GuestMiddlewre = "guest"

// AuthGuest using for ..
func AuthGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessAll := c.GetBool(AdminCanAccessAllKey)
		if accessAll {
			c.Next()
			return
		}

		userSessionName := viper.GetString("session.userID")
		helpers.StoreSessionString(c.Request, c.Writer, userSessionName, GuestMiddlewre)

		c.Next()
	}
}
