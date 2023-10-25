package service

import (
	"context"
	"errors"
	"tategoto/config/msg/errmsg"
	"tategoto/model"
	"tategoto/pkg/auth"
	"tategoto/pkg/ulid"
)

func (us *userService) RestoreUser(ctx context.Context, token string) (*model.User, error) {
	//jwtの検証
	userID, err := auth.VerifyUserJWT(token)
	if err != nil {
		return nil, err
	}

	user, err := us.ur.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) SignUp(ctx context.Context, user *model.User) (*model.User, error) {
	//IDの生成
	id, err := ulid.CreateULID()
	if err != nil {
		return nil, errors.New(errmsg.GenerateIDErr)
	}
	user.ID = id
	//メールアドレス重複チェック
	spUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return nil, err
	} else if spUser.Mail != "" {
		//重複エラー
		return nil, errors.New(errmsg.DuplicateMailErr)
	}

	//パスワード暗号化
	pw, err := auth.EncryptPassword(user.Password)
	if err != nil {
		//暗号化エラー
		return nil, errors.New(errmsg.EncryptionErr)
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
		return nil, "", errors.New(errmsg.IncorrectMailOrPasswordErr)
	}

	//パスワード比較
	err = auth.CompareHashAndPassword(spUser.Password, user.Password)
	if err != nil {
		//パスワード不一致エラー
		return nil, "", errors.New(errmsg.IncorrectMailOrPasswordErr)
	}

	//token作成
	token, err := auth.CreateUserJWT(spUser.ID)
	if err != nil {
		return nil, "", err
	}
	return spUser, token, nil
}
