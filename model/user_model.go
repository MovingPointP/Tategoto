package model

import "gorm.io/gorm"

// TODO: 文字列制限
type User struct {
	gorm.Model
	NickName   string `json:"nick_name"`
	UniqueName string `json:"unique_name"`
	Password   string `json:"password"`
}

var SampleUser = User{
	NickName: "John", UniqueName: "john01", Password: "3576dhuw",
}
