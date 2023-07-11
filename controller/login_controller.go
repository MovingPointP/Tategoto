package controller

import (
	"net/http"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := serviceInstance.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, "")
	}
}
