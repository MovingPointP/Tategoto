package config

import (
	"os"
	"strconv"
	"tategoto/model"

	"github.com/joho/godotenv"
)

var ServConf *model.ServerConfig
var SmtpConf *model.SmtpConfig

func InitAppConfig() {
	if err := godotenv.Load("../.env.local"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_HOUR"))
	ServConf = &model.ServerConfig{
		DSN:               os.Getenv("DSN"),
		PORT:              os.Getenv("SERVER_PORT"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
	}
}

func InitTestConfig() {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_HOUR"))
	ServConf = &model.ServerConfig{
		DSN:               os.Getenv("DSN"),
		PORT:              os.Getenv("SERVER_PORT"),
		ACCESS_TOKEN_HOUR: access_token_hour,
		SECRET_KEY:        os.Getenv("SECRET_KEY"),
	}
}

func InitSmtpConfig() {
	SmtpConf = &model.SmtpConfig{
		SMTP_SERVER:   os.Getenv("SMTP_SERVER"),
		SMTP_PORT:     os.Getenv("SMTP_PORT"),
		AUTH_ADDRESS:  os.Getenv("AUTH_ADDRESS"),
		AUTH_PASSWORD: os.Getenv("AUTH_PASSWORD"),
	}
}
