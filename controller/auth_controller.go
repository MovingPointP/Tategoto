package controller

import (
	"net/http"
	"strconv"
	"tategoto/auth"
	"tategoto/config"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

// Middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")

		//tokenが存在しない場合
		if err != nil {
			//loginにリダイレクト
			//c.Redirect(http.StatusMovedPermanently, "/login")
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		//jwtの検証
		userId, err := auth.VerifyUserJWT(token)

		//user取得
		user, err := serviceInstance.GetUserById(c, userId)
		if err != nil {
			//loginにリダイレクト
			//c.Redirect(http.StatusMovedPermanently, "/login")
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
		} else {
			//contextにセット
			c.Set("AuthorizedUser", user)
			c.Next() //これより前は事前処理、後は事後処理
		}
	}
}

func signup(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := serviceInstance.SignUp(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		//TODO: 情報の一部のみ持つuserをreturn
		ctx.JSON(http.StatusOK, &user)
	}
}

func login(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := serviceInstance.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		//TODO: 情報の一部のみ持つuserをreturn
		token, err := auth.CreateUserJWT(strconv.Itoa(int(user.ID)))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.SetCookie("token", token, config.Config.ACCESS_TOKEN_HOUR*3600, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, &user)
	}
}
