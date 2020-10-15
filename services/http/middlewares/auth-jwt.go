package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	header "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
)

// HeaderJwt is using
type HeaderJwt header.HeaderJwt

// AuthJwt using for ..
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := HeaderJwt{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		c.Next()
	}
}
