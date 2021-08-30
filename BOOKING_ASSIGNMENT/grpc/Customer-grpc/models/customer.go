package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id           int64          `gorm:"integer"`
	CustomerName string         `gorm:"type:varchar(200);not null"`
	Address      sql.NullString `gorm:"type:varchar(200)"`
	License      string         `gorm:"type:varchar(20)"`
	Phone        string         `gorm:"type:varchar(12)"`
	Email        string         `gorm:"type:varchar(200)"`
	Password     string         `gorm:"type:varchar(200)"`
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
type MessageChangePassword struct {
	Message string `gorm:"type:varchar(200)"`
}
