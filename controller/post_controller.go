package controller

import (
	"net/http"
	"tategoto/model"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context) {
	var post model.Post
	//postにバインド
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	spPost, err := serviceInstance.CreatePost(ctx, &post)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post": spPost})
}

func getPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	pid, err := funk.StringToUint(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	post, err := serviceInstance.GetPostById(ctx, pid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"post": post})
	}
}

func getPosts(ctx *gin.Context) {

	userId := ctx.Query("uid")
	uid, err := funk.StringToUint(userId)
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
