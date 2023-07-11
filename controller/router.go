package controller

import (
	"tategoto/repository"
	"tategoto/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var repositoryInstance repository.Repositorys
var serviceInstance service.Services

func GetRouter(db *gorm.DB) *gin.Engine {
	//create engine
	r := gin.Default()

	repositoryInstance = repository.New(*db)
	serviceInstance = service.New(repositoryInstance)

	//routing
	r.GET("/users/:name", getUsersByName)
	r.GET("/signup", createUser)

	return r
}
