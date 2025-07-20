package routes

import (
	"be-event/controllers"
	"github.com/gin-gonic/gin"
	"be-event/middlewares"

)

func InitRouter() *gin.Engine {
	r := gin.Default()

	eventController := NewEventController()
	authController := NewAuthController()

	ticketController := NewTicketController()
	RegisterEventRoutes(r, eventController)
	RegisterAuthRoutes(r, authController)
	RegisterTicketRoutes(r, ticketController)

	return r
}

func RegisterEventRoutes(r *gin.Engine, controller *controllers.EventController) {
	events := r.Group("/events")
	{
		events.GET("/api/list-events", controller.ListEvents)
		events.POST("/api/create-events",  middlewares.AuthMiddleware() ,controller.CreateEvent)

	}
}

func RegisterTicketRoutes(r *gin.Engine, controller *controllers.TicketController) {
	tickets := r.Group("/tickets")
	{
		// tickets.GET("/api/list-tickets", controller.ListTickets)
		tickets.POST("/api/create-ticket",middlewares.AuthMiddleware() ,controller.CreateTicket)
		// tickets.GET("/api/ticket/:id", controller.GetTicketByID)
		// tickets.PUT("/api/update-ticket/:id", controller.UpdateTicket)
		// tickets.DELETE("/api/delete-ticket/:id", controller.DeleteTicket)
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
