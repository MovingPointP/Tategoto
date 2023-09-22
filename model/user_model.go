package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID string `json:"ID" gorm:"primaryKey"`
	gorm.Model
	Mail     string `json:"mail" gorm:"unique;not null" validate:"email"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name"`
}
