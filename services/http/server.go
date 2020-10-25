package http

import (
	"strconv"
	"net/http"
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/controllers"
	"github.com/nugrohosam/gosampleapi/services/http/exceptions"
	"github.com/spf13/viper"
	"github.com/google/uuid"
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
	if viper.GetString("websocket.active") == "true" {
		fmt.Println("Hub connected")
		go HubConn.run()
	}

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

	// ws handler get
	if viper.GetString("websocket.active") == "true" {
		// load html example websocket chat
		Routes.LoadHTMLFiles("./views/websocket-chat.html")
		websocketChat := Routes.Group("/websocket-chat")
		{
			websocketChat.GET("/:roomID", func(ctx *gin.Context) {
				ctx.HTML(200, "websocket-chat.html", nil)
			})

			websocketChat.GET("/", func(ctx *gin.Context) {
				uuidClock := uuid.ClockSequence()
				roomID := strconv.Itoa(uuidClock)
				ctx.Redirect(http.StatusTemporaryRedirect, "/websocket-chat/" + roomID)
			})
		}
		
		Routes.GET("/ws/:roomID", func(ctx *gin.Context) {
			roomID := ctx.Param("roomID")
			ServerWS(ctx.Writer, ctx.Request, roomID)
		})
	}

	// v1
	v1 := Routes.Group("/v1")

	// v1/auth
	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.AuthHandlerLogin())
		auth.POST("/register", controllers.AuthHandlerRegister())
	}
}
