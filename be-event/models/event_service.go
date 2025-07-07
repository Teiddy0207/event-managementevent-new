package models

type EventService struct {
	ID        uint `gorm:"primaryKey"`
	EventID   uint
	ServiceID uint

	Event   Event   `gorm:"foreignKey:EventID"`
	Service Service `gorm:"foreignKey:ServiceID"`
}
