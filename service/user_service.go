package service

import (
	"context"
	"tategoto/model"
	"tategoto/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUsersByName(ctx context.Context, name string) ([]*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{ur: *ur}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) error {
	return us.ur.CreateUser(ctx, user)
}

func (us *userService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return us.ur.GetUserById(ctx, id)
}

func (us *userService) GetUsersByName(ctx context.Context, name string) ([]*model.User, error) {
	return us.ur.GetUsersByName(ctx, name)
}
