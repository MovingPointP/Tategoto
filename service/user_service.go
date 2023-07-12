package service

import (
	"context"
	"errors"
	"tategoto/config/msg"
	"tategoto/crypto"
	"tategoto/model"
	"tategoto/repository"
)

type UserService interface {
	SignUp(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUsersByName(ctx context.Context, name string) ([]*model.User, error)
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{ur: *ur}
}

func (us *userService) SignUp(ctx context.Context, user *model.User) error {
	//メールアドレス重複チェック
	receivedUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return err
	} else if receivedUser.Mail != "" {
		//重複エラー
		return errors.New(msg.DuplicateMailErr)
	}
	//パスワード暗号化
	pw, err := crypto.EncryptPassword(user.Password)
	if err != nil {
		//暗号化エラー
		return errors.New(msg.EncryptionErr)
	}
	user.Password = pw
	return us.ur.CreateUser(ctx, user)
}

func (us *userService) Login(ctx context.Context, user *model.User) error {
	receivedUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return err
	} else if receivedUser.Mail == "" {
		//メール非存在エラー
		return errors.New(msg.IncorrectMailOrPasswordErr)
	}
	err = crypto.CompareHashAndPassword(receivedUser.Password, user.Password)
	if err != nil {
		//パスワード不一致エラー
		return errors.New(msg.IncorrectMailOrPasswordErr)
	}
	return nil
}

func (us *userService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return us.ur.GetUserById(ctx, id)
}

func (us *userService) GetUsersByName(ctx context.Context, name string) ([]*model.User, error) {
	return us.ur.GetUsersByName(ctx, name)
}
