package models

import "time"

//Cities model
type Cities struct {
	CitiesID  uint64     `gorm:"column:cities_id" json:"cities_id"`
	Name      string     `gorm:"column:name" json:"name"`
	StateID   uint64     `gorm:"column:state_id" json:"state_id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
