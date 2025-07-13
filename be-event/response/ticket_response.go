package response

type TicketResponse struct {
	ID                uint    `json:"id"`
	EventID           uint    `json:"event_id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	QuantityAvailable int     `json:"quantity_available"`
}
