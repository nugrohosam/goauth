package http

import (
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/exceptions"
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
	// run hub connection for websocket
	go HubConn.run()

	Routes = gin.New()
	Routes.Use(exceptions.Recovery500())
	Routes.Static("/assets", "./assets")
	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// test-sentry
	Routes.GET("/test-sentry", func(ctx *gin.Context) {
		panic("make panic test")
	})

	// load html example websocket chat
	Routes.LoadHTMLFiles("./views/websocket-chat.html")
	Routes.GET("/websocket-chat", func(ctx *gin.Context) {
		ctx.HTML(200, "websocket.html", nil)
	})

	Routes.GET("/ws/:roomId", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")
		ServerWS(ctx.Writer, ctx.Request, roomId)
	})

	// v1
	v1 := Routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}
}
