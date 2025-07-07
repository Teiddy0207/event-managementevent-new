package services

import (
	"be-event/dto"
	"be-event/models"
	"be-event/repositories"
	"errors"
	"time"
)

type EventService struct {
	repo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(dto *dto.EventDTO) (*models.Event, error) {
	// Parse ngày
	eventDate, err := time.Parse("2006-01-02", dto.EventDate)
	if err != nil {
		return nil, errors.New("Ngày sự kiện không hợp lệ (YYYY-MM-DD)")
	}

	// Parse giờ bắt đầu
	startTime, err := time.Parse("15:04:05", dto.StartTime)
	if err != nil {
		return nil, errors.New("Giờ bắt đầu không hợp lệ (HH:mm:ss)")
	}

	// Parse giờ kết thúc
	endTime, err := time.Parse("15:04:05", dto.EndTime)
	if err != nil {
		return nil, errors.New("Giờ kết thúc không hợp lệ (HH:mm:ss)")
	}

	event := models.Event{
		Title:       dto.Title,
		Description: dto.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		EventTypeID: dto.EventTypeID,
		LocationID:  dto.LocationID,
		UserID:      dto.UserID,
		EventDate:   eventDate,
	}

	err = s.repo.CreateEvent(&event)
	if err != nil {
		return nil, err
	}

	if len(dto.ServiceIDs) > 0 {
		if err := s.repo.AttachServices(event.ID, dto.ServiceIDs); err != nil {
			return nil, err
		}
	}

	return &event, nil
}
