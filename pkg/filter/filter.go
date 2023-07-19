package filter

import (
	"tategoto/model"
	"time"

	"gorm.io/gorm"
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
	SocUser.CreatedAt = time.Time{}
	SocUser.UpdatedAt = time.Time{}
	SocUser.DeletedAt = gorm.DeletedAt{}
	SocUser.Mail = ""
	return SocUser
}
