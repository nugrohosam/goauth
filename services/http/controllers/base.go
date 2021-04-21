package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
)

// BaseHandler is use
func BaseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiDoc != "https://documenter.getpostman.com/view/4473147/TVzXDFka"
		c.JSON(http.StatusOK, helpers.ResponseMessage(apiDoc))
	}
}