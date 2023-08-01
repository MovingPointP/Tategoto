package service

import (
	"context"
	"errors"
	"tategoto/config/msg"
	"tategoto/model"
	"tategoto/pkg/auth"
)

func (us *userService) RestoreUser(ctx context.Context, token string) (*model.User, error) {
	//jwtの検証
	userId, err := auth.VerifyUserJWT(token)
	if err != nil {
		return nil, err
	}

	user, err := us.ur.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) SignUp(ctx context.Context, user *model.User) (*model.User, error) {
	//メールアドレス重複チェック
	spUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return nil, err
	} else if spUser.Mail != "" {
		//重複エラー
		return nil, errors.New(msg.DuplicateMailErr)
	}

	//パスワード暗号化
	pw, err := auth.EncryptPassword(user.Password)
	if err != nil {
		//暗号化エラー
		return nil, errors.New(msg.EncryptionErr)
	}
	user.Password = pw

	//ユーザー作成
	spUser, err = us.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return spUser, nil
}

// TODO: unique_nameでのLogin
func (us *userService) Login(ctx context.Context, user *model.User) (*model.User, string, error) {
	//ユーザー取得
	spUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return nil, "", err
	} else if spUser.Mail == "" {
		//メール非存在エラー
		return nil, "", errors.New(msg.IncorrectMailOrPasswordErr)
	}

	//パスワード比較
	err = auth.CompareHashAndPassword(spUser.Password, user.Password)
	if err != nil {
		//パスワード不一致エラー
		return nil, "", errors.New(msg.IncorrectMailOrPasswordErr)
	}

	//token作成
	token, err := auth.CreateUserJWT(spUser.ID)
	if err != nil {
		return nil, "", err
	}
	return spUser, token, nil
}
