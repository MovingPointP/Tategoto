package service

import (
	"tategoto/repository"
)

type Services interface {
	UserService
	PostService
}

type services struct {
	*userService
	*postService
}

func New(repo repository.Repositorys) Services {
	return &services{
		userService: &userService{ur: repo},
		postService: &postService{pr: repo},
	}
}
