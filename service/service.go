package service

import (
	"tategoto/repository"
)

type Services interface {
	UserService
}

type services struct {
	*userService
}

func New(repo repository.Repositorys) Services {
	return &services{
		userService: &userService{ur: repo},
	}
}
