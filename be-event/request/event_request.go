package request

type CreateEventRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
	EventTypeID uint   `json:"event_type_id" binding:"required"`
	LocationID  uint   `json:"location_id" binding:"required"`
	// UserID      uint   `json:"user_id"` // Tạm thời chưa có middleware
	EventDate   string `json:"event_date" binding:"required"`
	ServiceIDs  []uint `json:"service_ids"` // Danh sách dịch vụ đi kèm
}
