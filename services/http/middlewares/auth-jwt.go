package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	header "github.com/nugrohosam/gosampleapi/services/http/requests/header"
)

// Header is using
type Header header.HeaderJwt

// AuthJwt using for ..
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := Header{}
		if err := c.ShouldBindHeader(&header); err != nil {
			panic(http.StatusUnauthorized)
		}
	}
}
