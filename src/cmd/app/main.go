package main

import (
	"tategoto/src/config"
	"tategoto/src/connect"
	"tategoto/src/controller"
	"tategoto/src/model"
)

func main() {
	//.envの読み込み
	config.InitAppConfig()

	//DBに接続
	db := connect.GetConnection()
	defer connect.CloseConnection(db)
	//migration
	db.AutoMigrate(&model.User{}, &model.Post{})

	//GinのEngine取得
	router := controller.GetRouter(db)
	//server起動
	router.Run(config.Config.PORT)
}
