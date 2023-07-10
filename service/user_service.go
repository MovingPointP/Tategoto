package service

import (
	"context"
	"tategoto/model"
	"tategoto/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User)
	GetUserById(ctx context.Context, id string)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{ur: *ur}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) {
	us.ur.CreateUser(ctx, user)
}

func (us *userService) GetUserById(ctx context.Context, id string) {
	us.ur.GetUserById(ctx, id)
}
