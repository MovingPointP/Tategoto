package model

import "gorm.io/gorm"

//TODO:モデル数が増えたらファイル分割

// TODO: 文字列制限
// TODO: index キー制約
type User struct {
	gorm.Model
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	UniqueName  string `json:"unique_name"`
	DisplayName string `json:"display_name"`
}
