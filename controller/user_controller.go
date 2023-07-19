package controller

import (
	"net/http"
	"tategoto/model"
	"tategoto/pkg/filter"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

func getUsersByName(ctx *gin.Context) {
	name := ctx.Param("name")
	users, err := serviceInstance.GetUsersByName(ctx, name)
	users = funk.Map(users, func(user *model.User) *model.User {
		return filter.SocialUser(user)
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	}
}
