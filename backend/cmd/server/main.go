package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-technat/self-host-planning-poker/backend/internal/db"
	"github.com/the-technat/self-host-planning-poker/backend/internal/routes"
	"github.com/the-technat/self-host-planning-poker/backend/internal/sockets"
)

func main() {
	r := gin.Default()
	db.InitDB()
	sockets.InitSockets()
	routes.InitRoutes(r)
	r.Run()
}
