package model

import "gorm.io/gorm"

// TODO: 文字列制限
// TODO: index
type User struct {
	gorm.Model
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
}
