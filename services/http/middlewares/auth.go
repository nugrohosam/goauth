package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	validator "github.com/go-playground/validator/v10"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	header "github.com/nugrohosam/gosampleapi/services/http/requests/v1"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// AuthJwt using for ..
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header header.HeaderJwt
		c.BindHeader(&header)

		validate := validator.New()
		if err := validate.Struct(header); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusUnauthorized, helpers.ResponseErrValidation(fieldsErrors))
			c.Abort()
			return
		}

		token := strings.Replace(header.Authorization, "Bearer ", "", len(header.Authorization))
		if err := usecases.AuthorizationValidation(token); err != nil {
			c.JSON(http.StatusNotAcceptable, helpers.ResponseErr(err.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}
