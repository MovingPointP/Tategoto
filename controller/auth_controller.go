package controller

import (
	"net/http"
	"strconv"
	"tategoto/auth"
	"tategoto/config"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("token") //cookieからtokenの取得

		//tokenが存在しない場合
		if err != nil {
			ctx.Set("pastURI", ctx.Request.RequestURI)          //元のURIを保持
			ctx.Redirect(http.StatusMovedPermanently, "/login") //loginにリダイレクト
			ctx.Abort()
		}

		user, err := serviceInstance.RestoreUser(ctx, token)
		if err != nil {
			ctx.Set("pastURI", ctx.Request.RequestURI)          //元のURIを保持
			ctx.Redirect(http.StatusMovedPermanently, "/login") //loginにリダイレクト
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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	receivedUser, err := serviceInstance.SignUp(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &receivedUser)

}

func login(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	receivedUser, err := serviceInstance.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//token作成
	token, err := auth.CreateUserJWT(strconv.Itoa(int(receivedUser.ID)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("token", token, config.Config.ACCESS_TOKEN_HOUR*3600, "/", "localhost", false, true) //cookieにセット

	//元のURIを取得
	pastURI, ok := ctx.Get("pastURI")
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": &user, "redirect": pastURI.(string)})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": &user, "redirect": ""})
	}

}
