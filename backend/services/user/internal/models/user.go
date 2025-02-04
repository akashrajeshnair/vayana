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
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return nil
}
