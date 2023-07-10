package repository

import (
	"context"
	"tategoto/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
	GetUserById(ctx context.Context, id string)
}

type userRepository struct {
	db gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: *db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Create(user)
}

func (ur *userRepository) GetUserById(ctx context.Context, id string) {
}
