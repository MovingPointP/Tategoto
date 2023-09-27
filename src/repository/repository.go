package repository

import (
	"gorm.io/gorm"
)

type Repositorys interface {
	UserRepository
	PostRepository
}

type repositorys struct {
	*userRepository
	*postRepository
}

func New(db gorm.DB) Repositorys {
	return &repositorys{
		userRepository: &userRepository{db: db},
		postRepository: &postRepository{db: db},
	}
}
