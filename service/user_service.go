package service

import (
	"context"
	"tategoto/model"
	"tategoto/repository"
)

type UserService interface {
	RestoreUser(ctx context.Context, token string) (*model.User, error)
	SignUp(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (*model.User, string, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{ur: *ur}
}

func (us *userService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return us.ur.GetUserById(ctx, id)
}

func (us *userService) GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error) {

	return us.ur.GetUsers(ctx, userOption)
}
