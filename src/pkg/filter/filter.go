package filter

import (
	"tategoto/src/model"
)

// TODO:filterのフィールドの初期化忘れやfilterし忘れを防ぎたい

func PersonalUser(user *model.User) *model.User {
	perUser := user
	perUser.Password = ""
	return perUser
}

func SocialUser(user *model.User) *model.User {
	SocUser := user
	SocUser.Password = ""
	SocUser.Mail = ""
	return SocUser
}
