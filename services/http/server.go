package http

import (
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/exceptions"
	"github.com/nugrohosam/gosampleapi/services/http/middlewares"
	"github.com/spf13/viper"
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
	Routes.Use(exceptions.Recovery500())
	Routes.Static("/assets", "./assets")
	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// test-sentry
	Routes.GET("/test-sentry", func(c *gin.Context) {
		panic("make panic test")
	})

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
