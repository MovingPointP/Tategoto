package controller

import (
	"tategoto/model"
	"tategoto/service"

	"github.com/gin-gonic/gin"
)

func getUserById(serv service.Services) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		serv.GetUserById(ctx, id)
	}
}

func createUser(serv service.Services) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		serv.CreateUser(ctx, &model.SampleUser)
	}
}
