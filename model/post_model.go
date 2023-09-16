package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id" gorm:"not null"`
}
