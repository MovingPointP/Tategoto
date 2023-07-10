package controller

import (
	"net/http"
	"tategoto/model"
	"tategoto/service"

	"github.com/gin-gonic/gin"
)

func getUsersByName(serv service.Services) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		users, err := serv.GetUsersByName(ctx, name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, users)
		}
	}
}

func createUser(serv service.Services) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := serv.CreateUser(ctx, &model.SampleUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, "")
		}
	}
}
