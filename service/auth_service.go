package service

import (
	"errors"
	"tategoto/auth"
	"tategoto/config/msg"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

func (us *userService) RestoreUser(ctx *gin.Context, token string) (*model.User, error) {
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

func (us *userService) SignUp(ctx *gin.Context, user *model.User) (*model.User, error) {
	//メールアドレス重複チェック
	receivedUser, err := us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return nil, err
	} else if receivedUser.Mail != "" {
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
	if err := us.ur.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	//ユーザー取得
	receivedUser, err = us.ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return nil, err
	}

	return receivedUser, nil
}

// TODO: unique_nameでのLogin
func (us *userService) Login(ctx *gin.Context, user *model.User) (*model.User, error) {
	//ユーザー取得
	receivedUser, err := us.ur.GetUserPasswordByMail(ctx, user.Mail)
	if err != nil {
		return nil, err
	} else if receivedUser.Mail == "" {
		//メール非存在エラー
		return nil, errors.New(msg.IncorrectMailOrPasswordErr)
	}

	//パスワード比較
	err = auth.CompareHashAndPassword(receivedUser.Password, user.Password)
	if err != nil {
		//パスワード不一致エラー
		return nil, errors.New(msg.IncorrectMailOrPasswordErr)
	}
	receivedUser.Password = ""
	return receivedUser, nil
}
