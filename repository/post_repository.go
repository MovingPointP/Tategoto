package repository

import (
	"context"
	"tategoto/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	//Insert
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
	//Select
	GetPostByID(ctx context.Context, id uint) (*model.Post, error)
	GetPosts(ctx context.Context, postOption *model.Post) ([]*model.Post, error)
}

type postRepository struct {
	db gorm.DB
}

func (pr *postRepository) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	result := pr.db.Create(&post)
	return post, result.Error
}

func (pr *postRepository) GetPostByID(ctx context.Context, id uint) (*model.Post, error) {
	var post *model.Post
	result := pr.db.
		Where("id = ?", id).
		Where("deleted_at is null").
		Find(&post)
	return post, result.Error
}

func (pr *postRepository) GetPosts(ctx context.Context, postOption *model.Post) ([]*model.Post, error) {
	var posts []*model.Post
	chain := pr.db.Where("deleted_at is null")

	if postOption.UserID != 0 {
		chain.Where("user_id = ?", postOption.UserID)
	}

	result := chain.Find(&posts)
	return posts, result.Error
}
