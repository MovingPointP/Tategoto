package controller

import (
	"net/http"
	"strconv"
	"tategoto/config"
	"tategoto/config/msg"
	"tategoto/model"
	"tategoto/pkg/auth"
	"tategoto/pkg/filter"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//cookieからtokenの取得
		token, err := ctx.Cookie("token")

		//tokenが存在しない場合
		if err != nil {
			ctx.JSON(http.StatusSeeOther, gin.H{"message": msg.ShouldLoginErr, "path": ctx.Request.URL.Path})
			ctx.Abort()
			return
		}

		//Userの復元
		user, err := serviceInstance.RestoreUser(ctx, token)
		if err != nil {
			ctx.JSON(http.StatusSeeOther, gin.H{"message": msg.ShouldLoginErr, "path": ctx.Request.URL.Path})
			ctx.Abort()
		} else {
			ctx.Set("authorizedUser", user) //userを保持
			ctx.Next()                      //この行より前は事前処理、後は事後処理
		}
	}
}

func signup(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	spUser, err := serviceInstance.SignUp(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": filter.PersonalUser(spUser)})

}

func login(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	spUser, err := serviceInstance.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//token作成
	token, err := auth.CreateUserJWT(strconv.Itoa(int(spUser.ID)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//cookieにセット
	ctx.SetCookie("token", token, config.Config.ACCESS_TOKEN_HOUR*3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"user": filter.PersonalUser(spUser)})

}
