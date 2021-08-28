package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	Id            int64  `gorm:"primarykey"`
	Name          string `gorm:"type:varchar(200)"`
	From          string `gorm:"type:varchar(200)"`
	To            string `gorm:"type:varchar(20)"`
	Status        string `gorm:"type:varchar(12)"`
	Date          time.Time
	AvailableSlot int64 `gorm:"type:integer"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
type ListFlights struct {
	Flights []*Flight
}
