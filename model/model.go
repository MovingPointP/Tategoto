package model

import "gorm.io/gorm"

// TODO: 文字列制限
type User struct {
	gorm.Model
	NickName string `json:"nick_name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

var SampleUser = User{
	NickName: "John", Mail: "john01@fef.com", Password: "3576dhuw",
}
