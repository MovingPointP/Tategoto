package controller

import (
	"net/http"
	"tategoto/model"
	"tategoto/pkg/filter"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := serviceInstance.GetUserById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": filter.SocialUser(user)})
	}
}

func getUsers(ctx *gin.Context) {
	userOption := &model.User{
		Name: ctx.Query("name"),
	}
	users, err := serviceInstance.GetUsers(ctx, userOption)
	filteredUsers := funk.Map(users, func(user *model.User) *model.User {
		return filter.SocialUser(user)
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"users": filteredUsers})
	}
}
