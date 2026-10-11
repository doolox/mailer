package mail

import (
	"log"
	"net/smtp"
)

// SMTP server connection settings for Brevo (formerly Sendinblue).
// The remaining configuration lives in config.yaml.
const (
	smtpHost = "smtp-relay.brevo.com" // Brevo SMTP server host
	smtpPort = "587"                  // 587 (STARTTLS, recommended), 2525, or 465 (SSL/TLS)
)

func Start() {
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	sendTest(cfg)
}

func sendTest(cfg *Config) {
	addr := smtpHost + ":" + smtpPort
	auth := smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, smtpHost)

	msg := []byte("To: " + cfg.ToAddress + "\r\n" +
		"From: " + cfg.FromName + " <" + cfg.FromAddress + ">\r\n" +
		"Subject: Test\r\n" +
		"\r\n" +
		"hello world\r\n")

	if err := smtp.SendMail(addr, auth, cfg.FromAddress, []string{cfg.ToAddress}, msg); err != nil {
		log.Fatal("failed to send test email: ", err)
	}

	log.Println("test email sent successfully")
}
