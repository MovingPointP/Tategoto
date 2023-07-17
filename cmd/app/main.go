package main

import (
	"tategoto/config"
	"tategoto/connect"
	"tategoto/controller"
	"tategoto/model"
)

func main() {
	//.envの読み込み
	config.InitConfig()

	//DBに接続
	db := connect.GetConnection()
	defer connect.CloseConnection(db)
	//migration
	db.AutoMigrate(&model.User{})

	//GinのEngine取得
	router := controller.GetRouter(db)
	//server起動
	router.Run(config.Config.PORT)
}
