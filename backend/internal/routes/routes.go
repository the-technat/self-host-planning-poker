package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/the-technat/self-host-planning-poker/backend/internal/handlers"
	"github.com/the-technat/self-host-planning-poker/backend/internal/sockets"
)

func InitRoutes(r *gin.Engine) {
	appTitle := "Planning Poker"
	appBase := "/"
	if !strings.HasSuffix(appBase, "/") {
		appBase += "/"
	}

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"AppTitle": appTitle,
			"AppBase":  appBase,
		})
	})

	r.Static("/static", "./static")

	r.POST("/create", func(c *gin.Context) {
		handlers.CreateGame(c)
	})

	r.Any("/socket.io/", gin.WrapH(sockets.GetServeHandler()))

	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"AppTitle": appTitle,
			"AppBase":  appBase,
		})
	})
}
