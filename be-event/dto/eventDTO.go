package dto

type EventDTO struct {
	Title       string
	Description string
	StartTime   string
	EndTime     string
	EventDate   string `json:"event_date"`
	EventTypeID uint
	LocationID  uint
	UserID      uint
	ServiceIDs  []uint
}
