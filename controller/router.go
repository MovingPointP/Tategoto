package controller

import (
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
	r.GET("/users/:name", getUsersByName)
	r.POST("/signup", signup)

	return r
}
