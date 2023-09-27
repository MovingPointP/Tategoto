package model

import (
	"gorm.io/gorm"
)

type Post struct {
	ID string `json:"ID" gorm:"primaryKey"`
	gorm.Model
	Content string `json:"content"`
	UserID  string `json:"user_id" gorm:"not null"`
}
