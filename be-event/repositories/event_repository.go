package repositories

import (
	"be-event/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *models.Event) error
	AttachServices(eventID uint, serviceIDs []uint) error
	FindAllEvents(events *[]models.Event) error
}

type eventRepository struct {
	masterDB  *gorm.DB
	replicaDB *gorm.DB
}

func NewEventRepository(masterDB, replicaDB *gorm.DB) EventRepository {
	return &eventRepository{

		masterDB:  masterDB,
		replicaDB: replicaDB,
	}
}

func (r *eventRepository) CreateEvent(event *models.Event) error {
	return r.masterDB.Create(event).Error
}

func (r *eventRepository) AttachServices(eventID uint, serviceIDs []uint) error {
	var relations []models.EventService
	for _, sid := range serviceIDs {
		relations = append(relations, models.EventService{
			EventID:   eventID,
			ServiceID: sid,
		})
	}
	return r.masterDB.Create(&relations).Error
}

func (r *eventRepository) FindAllEvents(events *[]models.Event) error {

	return r.replicaDB.Preload("EventType").
		Preload("Location").
		Preload("Services").
		Preload("Tickets").
		Find(events).Error
}
