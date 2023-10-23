package services

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

func SendMail(to string, subject string, body string) {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587" // Typically 587 for TLS or 465 for SSL
	senderEmail := "salmannugas@gmail.com"
	senderPassword := "zngslspvujfbnerl"

	// Message.
	message := "To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Connect to the Gmail SMTP server.
	tlsConfig := &tls.Config{
		ServerName: smtpServer,
	}
	client, err := smtp.Dial(smtpServer + ":" + smtpPort)
	if err != nil {
		log.Fatal(err)
	}

	// Authenticate and send the email.
	if err = client.StartTLS(tlsConfig); err != nil {
		log.Fatal(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Fatal(err)
	}

	if err = client.Mail(senderEmail); err != nil {
		log.Fatal(err)
	}

	if err = client.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	messageWriter, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = messageWriter.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	err = messageWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	client.Quit()
}
