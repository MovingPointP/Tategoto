package apitest

import "tategoto/model"

type resUser struct {
	User *model.User
}

type resUsers struct {
	Users []*model.User
}

type resPost struct {
	Post *model.Post
}

type resPosts struct {
	Posts []*model.Post
}

type resFail struct {
	Message string
}
