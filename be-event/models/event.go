package models

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	EventTypeID uint      `json:"event_type_id"`
	LocationID  uint      `json:"location_id"`
	UserID      uint      `json:"user_id"`
	EventDate   time.Time `json:"event_date"` // Ngày diễn ra sự kiện
	// Quan hệ
	EventType EventType `gorm:"foreignKey:EventTypeID"`
	Location  Location  `gorm:"foreignKey:LocationID"`
	User      User      `gorm:"foreignKey:UserID"`
	Services  []Service `gorm:"many2many:event_services"`
	CreatedAt time.Time

}
