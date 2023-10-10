package controller

import (
	"net/http"
	"tategoto/config/msg/errmsg"
	"tategoto/model"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context) {
	pos, _ := ctx.Get("Post")
	post, ok := pos.(*model.Post)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": errmsg.PostBindErr})
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

	post, err := serviceInstance.GetPostByID(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if post.ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errmsg.NoDataErr})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

func getPosts(ctx *gin.Context) {

	userID := ctx.Query("uid")
	postOption := &model.Post{
		UserID: userID,
	}
	posts, err := serviceInstance.GetPosts(ctx, postOption)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(posts) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errmsg.NoDataErr})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}
