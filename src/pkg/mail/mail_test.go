package mail

import (
	"tategoto/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpMail(t *testing.T) {
	config.InitSmtpConfig()

	//test用メ―ルアドレス
	toMailAdress := "demyazuryo@nekosan.uk"
	err := SendSignUpMail(toMailAdress)
	assert.Equal(t, nil, err)
}
