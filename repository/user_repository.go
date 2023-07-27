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
	GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error)
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
	result := ur.db.
		Where("id = ?", id).
		Where("deleted_at is null").
		Find(&user)
	return user, result.Error
}

func (ur *userRepository) GetUserByMail(ctx context.Context, mail string) (*model.User, error) {
	var user *model.User
	result := ur.db.
		Where("mail = ?", mail).
		Where("deleted_at is null").
		Find(&user)
	return user, result.Error
}

func (ur *userRepository) GetUsers(ctx context.Context, userOption *model.User) ([]*model.User, error) {
	var users []*model.User
	chain := ur.db.Where("deleted_at is null")

	if userOption.Name != "" {
		chain.Where("name = ?", userOption.Name)
	}

	result := chain.Find(&users)
	return users, result.Error
}
