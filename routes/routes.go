package routes

import (
	"antero.com/event_booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents) // GET,POST,PUT,PATCH,DELETE
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	server.POST("/signup", signup)
	server.POST("/login", login)

}
