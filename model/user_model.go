package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Mail     string `json:"mail" gorm:"unique;not null" validate:"email"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name"`
}
