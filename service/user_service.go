package service

import (
	"tategoto/model"
	"tategoto/repository"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RestoreUser(ctx *gin.Context, token string) (*model.User, error)
	SignUp(ctx *gin.Context, user *model.User) (*model.User, error)
	Login(ctx *gin.Context, user *model.User) (*model.User, error)
	GetUserById(ctx *gin.Context, id string) (*model.User, error)
	GetUsersByName(ctx *gin.Context, name string) ([]*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{ur: *ur}
}

func (us *userService) GetUserById(ctx *gin.Context, id string) (*model.User, error) {
	return us.ur.GetUserById(ctx, id)
}

func (us *userService) GetUsersByName(ctx *gin.Context, name string) ([]*model.User, error) {
	return us.ur.GetUsersByName(ctx, name)
}
