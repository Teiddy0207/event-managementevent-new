package services

import (
	"be-event/dto"
	"be-event/models"
	"be-event/repositories"
	"errors"
)

type TicketService struct {
	repo repositories.TicketRepository

	//repo repository.TicketRepository
}

func NewTicketService(repo repositories.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(dtos []dto.TicketDTO) ([]models.Ticket, error) {
	if len(dtos) == 0 {
		return nil, errors.New("Không có vé nào được gửi lên")
	}

	eventID := dtos[0].EventID
	event, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return nil, errors.New("Sự kiện không tồn tại")
	}
	if event.EventTypeID != 1 {
		return nil, errors.New("Chỉ tạo vé cho sự kiện admin (event_type_id = 1)")
	}

	// Kiểm tra các vé có cùng event_id
	for _, dto := range dtos {
		if dto.EventID != eventID {
			return nil, errors.New("Tất cả vé phải cùng một sự kiện")
		}
	}

	var tickets []models.Ticket
	for _, dto := range dtos {
		tickets = append(tickets, models.Ticket{
			EventID:           dto.EventID,
			Name:              dto.Name,
			Price:             dto.Price,
			QuantityAvailable: dto.QuantityAvailable,
		})
	}

	if err := s.repo.Create(tickets); err != nil {
		return nil, err
	}

	return tickets, nil

}
