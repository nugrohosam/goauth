package http

import (
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
)

// Routes ...
var Routes *gin.Engine

// Serve using for listen to specific port
func Serve() error {
	Prepare()

	port := viper.GetString("app.port")
	if err := Routes.Run(":" + port); err != nil {
		return err
	}

	return nil
}

// Prepare ...
func Prepare() {
	Routes = gin.New()
	Routes.Use(gin.Logger())
	Routes.Use(gin.Recovery())

	// v1
	v1 := Routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}

	// v1/hello-with-middleware
	helloWithMiddleware := v1.Group("/hello-with-middleware")
	helloWithMiddleware.Use(middlewares.AuthJwt())
	{
		helloWithMiddleware.GET("/", controllers.HelloWorldHandler())
	}

	helloWithoutMiddleware := v1.Group("/hello-without-middleware")
	{
		helloWithoutMiddleware.GET("/", controllers.HelloWorldHandler())
	}
}
