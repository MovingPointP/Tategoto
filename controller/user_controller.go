package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsersByName(ctx *gin.Context) {
	name := ctx.Param("name")
	users, err := serviceInstance.GetUsersByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}
