package service

import (
	"context"
	"errors"
	"tategoto/auth"
	"tategoto/config/msg"
	"tategoto/model"
)

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
	pw, err := auth.EncryptPassword(user.Password)
	if err != nil {
		//暗号化エラー
		return errors.New(msg.EncryptionErr)
	}
	user.Password = pw
	return us.ur.CreateUser(ctx, user)
}

// TODO: unique_nameでのLogin
func (us *userService) Login(ctx context.Context, user *model.User) error {
	receivedUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return err
	} else if receivedUser.Mail == "" {
		//メール非存在エラー
		return errors.New(msg.IncorrectMailOrPasswordErr)
	}
	err = auth.CompareHashAndPassword(receivedUser.Password, user.Password)
	if err != nil {
		//パスワード不一致エラー
		return errors.New(msg.IncorrectMailOrPasswordErr)
	}
	return nil
}
