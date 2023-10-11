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
		Dsn:             os.Getenv("DSN"),
		Port:            os.Getenv("SERVER_PORT"),
		AccessTokenHour: access_token_hour,
		SecretKey:       os.Getenv("SECRET_KEY"),
	}
}

func InitTestConfig() {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	access_token_hour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_HOUR"))
	ServConf = &model.ServerConfig{
		Dsn:             os.Getenv("DSN"),
		Port:            os.Getenv("SERVER_PORT"),
		AccessTokenHour: access_token_hour,
		SecretKey:       os.Getenv("SECRET_KEY"),
	}
}

func InitSmtpConfig() {
	SmtpConf = &model.SmtpConfig{
		SmtpServer:   os.Getenv("SMTP_SERVER"),
		SmtpPort:     os.Getenv("SMTP_PORT"),
		AuthAddress:  os.Getenv("AUTH_ADDRESS"),
		AuthPassword: os.Getenv("AUTH_PASSWORD"),
	}
}
