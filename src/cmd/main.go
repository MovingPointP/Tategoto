package main

import (
	"tategoto/config"
	"tategoto/connect"
	"tategoto/controller"
	"tategoto/model"
)

func main() {
	//.envの読み込み
	config.InitAppConfig()
	config.InitSmtpConfig()

	//DBに接続
	db := connect.GetConnection()
	defer connect.CloseConnection(db)
	//migration
	db.AutoMigrate(&model.User{}, &model.Post{})

	//GinのEngine取得
	router := controller.GetRouter(db)
	//server起動
	router.Run(config.ServConf.Port)
}
