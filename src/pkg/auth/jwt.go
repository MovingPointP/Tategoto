package auth

import (
	"errors"
	"tategoto/config"
	"tategoto/config/msg/errmsg"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateUserJWT(userID string) (string, error) {
	//ペイロード
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.ServConf.AccessTokenHour)).Unix(), //トークン期限
	}

	//token生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//tokenに署名を付与
	tokenString, err := token.SignedString([]byte(config.ServConf.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyUserJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}
		return []byte(config.ServConf.SecretKey), nil
	})

	if err != nil {
		return "", errors.New(errmsg.VerifyTokenErr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	//token検証エラー
	if !ok || !token.Valid {
		return "", errors.New(errmsg.VerifyTokenErr)
	}

	//token期限切れ
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return "", errors.New(errmsg.ExpiredTokenErr)
	}

	id := claims["user_id"].(string)

	return id, nil

}
