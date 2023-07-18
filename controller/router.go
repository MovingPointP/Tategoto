package controller

import (
	"net/http"
	"tategoto/repository"
	"tategoto/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var serviceInstance service.Services

func GetRouter(db *gorm.DB) *gin.Engine {
	//engine作成
	r := gin.Default()

	//instance作成
	repositoryInstance := repository.New(*db)
	serviceInstance = service.New(repositoryInstance)

	//routing
	r.POST("/signup", signup)
	r.POST("/login", login)

	api := r.Group("/api")
	api.Use(AuthMiddleware()) //事前・事後処理
	{
		api.GET("/users/:name", getUsersByName)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404:NOT FOUND")
	})

	return r
}
