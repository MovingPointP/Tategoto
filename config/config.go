package config

import (
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
	if err := godotenv.Load("cmd/app/.env"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("access_token_hour"))
	Config = configList{
		DSN:               os.Getenv("dsn"),
		PORT:              os.Getenv("server_port"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("secret_key"),
	}
}

func InitTestConfig() {
	if err := godotenv.Load("../apitest/.env"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("access_token_hour"))
	Config = configList{
		DSN:               os.Getenv("dsn"),
		PORT:              os.Getenv("server_port"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("secret_key"),
	}
}
