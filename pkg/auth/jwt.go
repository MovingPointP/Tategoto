package auth

import (
	"errors"
	"tategoto/config"
	"tategoto/config/msg"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateUserJWT(userID string) (string, error) {
	//ペイロード
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.Config.ACCESS_TOKEN_HOUR)).Unix(), //トークン期限
	}

	//token生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//tokenに署名を付与
	tokenString, err := token.SignedString([]byte(config.Config.SECRET_KEY))
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
		return []byte(config.Config.SECRET_KEY), nil
	})

	if err != nil {
		return "", errors.New(msg.VerifyTokenErr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	//token検証エラー
	if !ok || !token.Valid {
		return "", errors.New(msg.VerifyTokenErr)
	}

	//token期限切れ
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return "", errors.New(msg.ExpiredTokenErr)
	}

	id := claims["user_id"].(string)

	return id, nil

}
