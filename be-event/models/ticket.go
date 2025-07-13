package models

type Ticket struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	EventID           uint    `json:"event_id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	QuantityAvailable int     `json:"quantity_available"`

	Event Event `gorm:"foreignKey:EventID"`
}
