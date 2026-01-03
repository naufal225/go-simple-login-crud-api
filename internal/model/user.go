package model

import "time"

type User struct {
	ID           string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(120);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:text;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

Items []Item `gorm:"foreignId:UserID"`