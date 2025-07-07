package models

type Service struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`

	EventServices []EventService `gorm:"foreignKey:ServiceID"`
}
