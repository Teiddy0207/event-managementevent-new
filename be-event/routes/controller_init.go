package routes

import (
	"be-event/config"
	"be-event/controllers"
	"be-event/repositories"
	"be-event/services"
)

func NewEventController() *controllers.EventController {
	repo := repositories.NewEventRepository(config.DBMaster, config.DBReplica)
	service := services.NewEventService(repo)
	return controllers.NewEventController(service)
}

func NewAuthController() *controllers.AuthController {
	repo := repositories.NewAuthRepository(config.DBMaster, config.DBReplica)
	service := services.NewAuthService(repo)
	return controllers.NewAuthController(service)
}

func NewTicketController() *controllers.TicketController {
	repo := repositories.NewTicketRepository(config.DBMaster, config.DBReplica)
	service := services.NewTicketService(repo)
	return controllers.NewTicketController(service)
}
