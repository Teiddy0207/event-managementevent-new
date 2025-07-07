package routes

import (
	"be-event/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	eventController := NewEventController()
	authController := NewAuthController()
	RegisterEventRoutes(r, eventController)
	RegisterAuthRoutes(r, authController)
	return r
}

func RegisterEventRoutes(r *gin.Engine, controller *controllers.EventController) {
	events := r.Group("/events")
	{
		events.GET("/api/list-events", controller.ListEvents)
		events.POST("/api/create-events", controller.CreateEvent)
	}
}

func RegisterAuthRoutes(r *gin.Engine, controller *controllers.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.POST("/logout", controller.Logout)
	}
}
