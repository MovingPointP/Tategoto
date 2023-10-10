package mail

import (
	"fmt"
	"net/smtp"
	"tategoto/config"
	"tategoto/config/msg/cmnmsg"
)

func SendSignUpMail(mailAddress string) error {
	toMailAdress := []string{mailAddress}

	subject := cmnmsg.SignUpSubject
	body := cmnmsg.SignUpBody
	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", mailAddress, subject, body))
	auth := smtp.CRAMMD5Auth(config.SmtpConf.AUTH_ADDRESS, config.SmtpConf.AUTH_PASSWORD)

	if err := smtp.SendMail(fmt.Sprintf("%s:%s", config.SmtpConf.SMTP_SERVER, config.SmtpConf.SMTP_PORT), auth, config.SmtpConf.AUTH_ADDRESS, toMailAdress, message); err != nil {
		return err
	}
	return nil
}
