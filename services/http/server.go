package http

import (
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/spf13/viper"

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

	// initial roles
	adminRole := viper.GetString("config.role.admin")

	Routes = gin.New()
	// Routes.Use(exceptions.Recovery500())
	Routes.Static("/assets", "./assets")
	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// test-sentry
	Routes.GET("/test-sentry", func(ctx *gin.Context) {
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

	role := v1.Group("/role")
	role.Use(middlewares.AuthJwt()).Use(middlewares.CanAccessBy(
		[]string{
			adminRole,
		},
	))
	{
		role.POST("/", controllers.RoleHandlerCreate())
		role.PUT("/:id", controllers.RoleHandlerUpdate())
		role.DELETE("/:id", controllers.RoleHandlerDelete())
	}

}
