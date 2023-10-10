package model

type ServerConfig struct {
	DSN               string
	PORT              string
	ACCESS_TOKEN_HOUR int
	SECRET_KEY        string
}

type SmtpConfig struct {
	SMTP_SERVER   string
	SMTP_PORT     string
	AUTH_ADDRESS  string
	AUTH_PASSWORD string
}
