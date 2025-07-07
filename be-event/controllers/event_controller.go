package controllers

import (
	"be-event/dto"
	"be-event/request"
	"be-event/response"
	"be-event/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct {
	service *services.EventService
}

func NewEventController(service *services.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) ListEvents(ctx *gin.Context) {
	events, err := c.service.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách sự kiện"})
		return
	}

	var res []response.EventResponse
	for _, e := range events {
		res = append(res, response.NewEventResponse(&e))
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var req request.CreateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Tạo DTO
	eventDTO := &dto.EventDTO{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		EventTypeID: req.EventTypeID,
		EventDate:   req.EventDate,
		LocationID:  req.LocationID,
		UserID:      req.UserID,
		ServiceIDs:  req.ServiceIDs,
	}

	// Gọi service
	event, err := c.service.CreateEvent(eventDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response.NewEventResponse(event))

}
