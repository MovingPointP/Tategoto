package service

import (
	"context"
	"tategoto/model"
	"tategoto/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
	GetPostByID(ctx context.Context, id string) (*model.Post, error)
	GetPosts(ctx context.Context, postOption *model.Post) ([]*model.Post, error)
}

type postService struct {
	pr repository.PostRepository
}

func (ps *postService) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	spPost, err := ps.pr.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return spPost, nil
}

func (ps *postService) GetPostByID(ctx context.Context, id string) (*model.Post, error) {

	return ps.pr.GetPostByID(ctx, id)
}

func (ps *postService) GetPosts(ctx context.Context, postOption *model.Post) ([]*model.Post, error) {

	return ps.pr.GetPosts(ctx, postOption)
}
