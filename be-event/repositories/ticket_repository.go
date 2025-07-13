package repositories

import (
	"be-event/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket []models.Ticket) error
	GetEventByID(eventID uint) (*models.Event, error)
}

type ticketRepository struct {
	masterDB  *gorm.DB
	replicaDB *gorm.DB
}

func NewTicketRepository(masterDB, replicaDB *gorm.DB) TicketRepository {
	return &ticketRepository{

		masterDB:  masterDB,
		replicaDB: replicaDB,
	}
}

func (r *ticketRepository) Create(ticket []models.Ticket) error {
	return r.masterDB.Create(&ticket).Error
}

func (r *ticketRepository) GetEventByID(eventID uint) (*models.Event, error) {
	var event models.Event
	if err := r.replicaDB.First(&event, eventID).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
