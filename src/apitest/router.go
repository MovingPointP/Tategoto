package apitest

import (
	"tategoto/config"
	"tategoto/connect"
	"tategoto/controller"
	"tategoto/model"
	"tategoto/repository"
	"tategoto/service"

	"github.com/gin-gonic/gin"
)

var serviceInstance service.Services

func NewRouter() *gin.Engine {
	//.envの読み込み
	config.InitTestConfig()

	//DBに接続
	db := connect.GetConnection()
	//defer connect.CloseConnection(db)

	//instance作成
	repositoryInstance := repository.New(*db)
	serviceInstance = service.New(repositoryInstance)

	//tableの削除
	db.Migrator().DropTable(&model.User{}, &model.Post{})
	//migration
	db.AutoMigrate(&model.User{}, &model.Post{})
	//GinのEngine取得
	router := controller.GetRouter(db)
	return router
}
