package main

import (
	"antero.com/event_booking/db"
	"antero.com/event_booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8090")

}
