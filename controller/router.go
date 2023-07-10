package controller

import (
	"tategoto/repository"
	"tategoto/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB) *gin.Engine {
	//create engine
	r := gin.Default()

	repository := repository.New(*db)
	service := service.New(repository)

	//routing
	r.GET("/users/:id", getUserById(service))
	r.GET("/signin", createUser(service))

	return r
}
