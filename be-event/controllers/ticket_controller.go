package controllers

import (
	"be-event/dto"
	"be-event/request"
	"be-event/response"
	"be-event/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	service *services.TicketService
}

func NewTicketController(service *services.TicketService) *TicketController {
	return &TicketController{service: service}
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var reqs []request.CreateTicketRequest
	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	var dtos []dto.TicketDTO
	for _, r := range reqs {
		dtos = append(dtos, dto.TicketDTO{
			EventID:           r.EventID,
			Name:              r.Name,
			Price:             r.Price,
			QuantityAvailable: r.QuantityAvailable,
		})
	}

	tickets, err := c.service.CreateTicket(dtos)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res []response.TicketResponse
	for _, t := range tickets {
		res = append(res, response.TicketResponse{
			ID:                t.ID,
			EventID:           t.EventID,
			Name:              t.Name,
			Price:             t.Price,
			QuantityAvailable: t.QuantityAvailable,
		})
	}

	ctx.JSON(http.StatusOK, res)
}
