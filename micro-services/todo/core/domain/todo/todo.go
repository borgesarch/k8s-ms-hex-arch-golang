package domain

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Description string `gorm:"size:255;not null"`
	Completed   bool
}
