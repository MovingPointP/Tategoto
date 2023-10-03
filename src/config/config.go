package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configList struct {
	DSN               string
	PORT              string
	ACCESS_TOKEN_HOUR int
	SECRET_KEY        string
}

var Config configList

func InitAppConfig() {
	// if err := godotenv.Load("cmd/app/.env"); err != nil {
	// 	panic(err)
	// }
	access_token_hour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_HOUR"))
	Config = configList{
		DSN:               os.Getenv("DSN"),
		PORT:              os.Getenv("SERVER_PORT"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
	}
	fmt.Println(Config.DSN)
}

func InitTestConfig() {
	if err := godotenv.Load("../apitest/.env"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_HOUR"))
	Config = configList{
		DSN:               os.Getenv("DSN"),
		PORT:              os.Getenv("SERVER_PORT"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
	}
}
