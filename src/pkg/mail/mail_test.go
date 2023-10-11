package mail

import (
	"tategoto/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSignUpMail(t *testing.T) {
	if err := godotenv.Load("../../../.env.local"); err != nil {
		panic(err)
	}
	config.InitSmtpConfig()

	//test用メ―ルアドレス
	toMailAdress := "demyazuryo@nekosan.uk"
	err := SendSignUpMail(toMailAdress)
	assert.Equal(t, nil, err)
}
