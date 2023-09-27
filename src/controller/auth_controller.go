package controller

import (
	"net/http"
	"tategoto/src/config"
	"tategoto/src/config/msg"
	"tategoto/src/model"
	"tategoto/src/pkg/filter"
	"tategoto/src/pkg/ulid"

	"github.com/gin-gonic/gin"
)

// tokenチェック
func tokenRequired() gin.HandlerFunc {
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
			return
		}

		ctx.Set("AuthorizedUser", user) //userを保持
		ctx.Next()                      //この行より前は事前処理、後は事後処理
	}
}

// tokenとpostのuserID比較
func compareTokenAndPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUser, _ := ctx.Get("AuthorizedUser")
		authorizedUser, ok := authUser.(*model.User)
		if !ok {
			ctx.JSON(http.StatusSeeOther, gin.H{"message": msg.ShouldLoginErr, "path": ctx.Request.URL.Path})
			ctx.Abort()
			return
		}

		var post model.Post
		//postにバインド
		if err := ctx.ShouldBindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": msg.PostBindErr})
			ctx.Abort()
			return
		}

		id, err := ulid.CreateULID()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": msg.GenerateIDErr})
			ctx.Abort()
			return
		}
		post.ID = id

		if authorizedUser.ID != post.UserID {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": msg.IncorrectUserIDErr, "path": ctx.Request.URL.Path})
			ctx.Abort()
			return
		}

		ctx.Set("Post", &post)
		ctx.Next()
	}
}

func signup(ctx *gin.Context) {
	var user model.User
	//userにバインド
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := ulid.CreateULID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": msg.GenerateIDErr})
		ctx.Abort()
		return
	}
	user.ID = id

	spUser, err := serviceInstance.SignUp(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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

	id, err := ulid.CreateULID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": msg.GenerateIDErr})
		ctx.Abort()
		return
	}
	user.ID = id

	spUser, token, err := serviceInstance.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//cookieにセット
	ctx.SetCookie("token", token, config.Config.ACCESS_TOKEN_HOUR*3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"user": filter.PersonalUser(spUser)})
}
