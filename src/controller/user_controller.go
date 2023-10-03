package controller

import (
	"net/http"
	"tategoto/config/msg"
	"tategoto/model"
	"tategoto/pkg/filter"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

func getUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := serviceInstance.GetUserByID(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if user.ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": msg.NoDataErr})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": filter.SocialUser(user)})
}

func getUsers(ctx *gin.Context) {
	userOption := &model.User{
		Name: ctx.Query("name"),
	}
	users, err := serviceInstance.GetUsers(ctx, userOption)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": msg.NoDataErr})
		return
	}

	var filteredUsers []*model.User
	if len(users) != 0 {
		filteredUsers = funk.Map(users, func(user *model.User) *model.User {
			return filter.SocialUser(user)
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"users": filteredUsers})
}
