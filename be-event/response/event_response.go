package response

import "be-event/models"

type EventResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EventType   string `json:"event_type"`
	Location    string `json:"location"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	EventDate string `json:"event_date"`
}

func NewEventResponse(event *models.Event) EventResponse {
	layout := "2006-01-02 15:04:05"

	return EventResponse{
		ID:          event.ID,
		Name:        event.Title,
		Description: event.Description,
		EventType:   event.EventType.Name,
		Location:    event.Location.Name,

		StartTime: event.StartTime.Format(layout),
		EndTime:   event.EndTime.Format(layout),
		EventDate:   event.EventDate.Format("2006-01-02"),
	}
}
