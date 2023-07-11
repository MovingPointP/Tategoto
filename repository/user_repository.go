package repository

import (
	"context"
	"tategoto/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	//Insert
	CreateUser(ctx context.Context, user *model.User) error
	//Select
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByMail(ctx context.Context, mail string) (*model.User, error)
	GetUsersByName(ctx context.Context, name string) ([]*model.User, error)
}

type userRepository struct {
	db gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: *db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := ur.db.Create(user)
	return result.Error
}

func (ur *userRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	result := ur.db.Find(&user, "id = ?", id)
	return user, result.Error
}

func (ur *userRepository) GetUserByMail(ctx context.Context, mail string) (*model.User, error) {
	var user *model.User
	result := ur.db.Find(&user, "mail = ?", mail)
	return user, result.Error
}

func (ur *userRepository) GetUsersByName(ctx context.Context, name string) ([]*model.User, error) {
	var users []*model.User
	result := ur.db.
		Select("id", "nick_name", "mail").
		Where("nick_name = ?", name).
		Find(&users)
	return users, result.Error
}
