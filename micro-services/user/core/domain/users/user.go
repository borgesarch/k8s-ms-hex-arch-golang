package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
}
