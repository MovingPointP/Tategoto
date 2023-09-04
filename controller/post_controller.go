package controller

import (
	"net/http"
	"tategoto/config/msg"
	"tategoto/model"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context) {
	pos, _ := ctx.Get("Post")
	post, ok := pos.(*model.Post)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": msg.PostBindErr})
		return
	}

	spPost, err := serviceInstance.CreatePost(ctx, post)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post": spPost})
}

func getPostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := funk.StringToUint(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	post, err := serviceInstance.GetPostByID(ctx, pid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"post": post})
	}
}

func getPosts(ctx *gin.Context) {

	userID := ctx.Query("uid")
	uid, err := funk.StringToUint(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	postOption := &model.Post{
		UserID: uid,
	}
	posts, err := serviceInstance.GetPosts(ctx, postOption)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"posts": posts})
	}
}
