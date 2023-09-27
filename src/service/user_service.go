package service

import (
	"context"
	"tategoto/src/model"
	"tategoto/src/repository"
)

type UserService interface {
	RestoreUser(ctx context.Context, token string) (*model.User, error)
	SignUp(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (*model.User, string, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func (us *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {

	return us.ur.GetUserByID(ctx, id)
}

func (us *userService) GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error) {

	return us.ur.GetUsers(ctx, userOption)
}
