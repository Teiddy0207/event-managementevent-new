package models

type EventType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
