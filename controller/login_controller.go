package controller

import (
	"net/http"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

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
		ctx.JSON(http.StatusOK, "success")
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
		//TODO: tokenをreturn
		ctx.JSON(http.StatusOK, "success")
	}
}
