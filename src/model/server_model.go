package model

type ServerConfig struct {
	Dsn             string
	Port            string
	AccessTokenHour int
	SecretKey       string
}

type SmtpConfig struct {
	SmtpServer   string
	SmtpPort     string
	AuthAddress  string
	AuthPassword string
}
