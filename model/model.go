package model

import "gorm.io/gorm"

//TODO:モデル数が増えたらファイル分割

// TODO: 文字列制限
// TODO: index キー制約
type User struct {
	gorm.Model
	Mail     string `json:"mail" gorm:"unique;not null" validate:"email"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name"`
}
}
