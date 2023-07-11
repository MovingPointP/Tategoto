package controller

import (
	"net/http"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

func getUsersByName(ctx *gin.Context) {
	name := ctx.Param("name")
	users, err := serviceInstance.GetUsersByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func createUser(ctx *gin.Context) {
	err := serviceInstance.CreateUser(ctx, &model.SampleUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, "")
	}
}
