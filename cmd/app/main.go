package main

import (
	"tategoto/config"
	"tategoto/connection"
	"tategoto/controller"
	"tategoto/model"
)

func main() {
	//.envの読み込み
	config.InitConfig()

	//DBに接続
	db := connection.GetConnection()
	defer connection.CloseConnection(db)
	//migration
	db.AutoMigrate(&model.User{})

	//GinのEngine取得
	router := controller.GetRouter(db)
	//server起動
	router.Run(config.Config.PORT)
}
