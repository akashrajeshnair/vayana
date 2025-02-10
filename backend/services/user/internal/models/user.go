package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primary_key;type:uuid"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GoogleID      string `gorm:"type:varchar(255);unique"`
	AuthProvider  string `gorm:"type:varchar(50)"`
	EmailVerified bool   `gorm:"default:false"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	if user.AuthProvider == "" {
		user.AuthProvider = "email"
	}

	if user.AuthProvider == "google" {
		user.EmailVerified = true
	}

	return nil
}
