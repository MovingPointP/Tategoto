package controller

import (
	"net/http"
	"tategoto/config/msg"
	"tategoto/model"
	"tategoto/pkg/funk"

	"github.com/gin-gonic/gin"
)

// tokenとpostのuserID比較
func CompareTokenAndPost(ctx *gin.Context, post *model.Post) bool {
	authUser, _ := ctx.Get("AuthorizedUser")
	authorizedUser, ok := authUser.(*model.User)
	if !ok {
		ctx.JSON(http.StatusSeeOther, gin.H{"message": msg.ShouldLoginErr, "path": ctx.Request.URL.Path})
		return false
	}
	return authorizedUser.ID == post.UserID
}

func createPost(ctx *gin.Context) {
	var post model.Post
	//postにバインド
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !CompareTokenAndPost(ctx, &post) {
		ctx.JSON(http.StatusSeeOther, gin.H{"message": msg.ShouldLoginErr, "path": ctx.Request.URL.Path})
		return
	}

	spPost, err := serviceInstance.CreatePost(ctx, &post)

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

func getPostsByUID(ctx *gin.Context) {

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
