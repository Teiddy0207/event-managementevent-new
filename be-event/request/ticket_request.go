package request

type CreateTicketRequest struct {
	EventID           uint    `json:"event_id" binding:"required"`
	Name              string  `json:"name" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
	QuantityAvailable int     `json:"quantity_available" binding:"required"`
}
