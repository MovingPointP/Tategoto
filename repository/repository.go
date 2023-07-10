package repository

import (
	"gorm.io/gorm"
)

type Repositorys interface {
	UserRepository
}

type repositorys struct {
	*userRepository
}

func New(db gorm.DB) Repositorys {
	return &repositorys{
		userRepository: &userRepository{db: db},
	}
}
