package model

import (
	"time"
	"github.com/naufal225/go-simple-login-crud-api/model"
)
type Item struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    string `gorm:"type:uuid;not null;index"`
	Name      string `gorm:"type:varchar(120);not null"`
	SKU       string `gorm:"type:varchar(64);uniqueIndex"`
	Price     int    `gorm:"not null;default:0"`
	Stock     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

User User `gorm:"foreignId:UserID"`