package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID      string `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	UserID  string `json:"user_id" gorm:"not null"`
}
