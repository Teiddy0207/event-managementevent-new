package models

type Location struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	CityID  uint   `json:"city_id"`
	Address string `json:"address"`

	City City `gorm:"foreignKey:CityID"`
}
