package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mail     string `json:"mail" gorm:"unique;not null" validate:"email"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name"`
}

type Post struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id" gorm:"not null"`
}
