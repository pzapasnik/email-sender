package email

import (
	"crypto/tls"
	"log/slog"
	"net/smtp"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (e *Service) Send() error {
	//SMTP conn
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "",
	}

	// Create connection
	conn, err := tls.Dial("tcp", "server189430.nazwa.pl:465", tlsConfig)
	if err != nil {
		slog.Error("Failed to connect to the server", "error", err)
		return err
	}

	slog.Info("Connected to the server")

	smtpClient, err := smtp.NewClient(conn, "server189430.nazwa.pl")
	if err != nil {
		slog.Error("Failed to create client", "error", err)
		return err
	}
	defer smtpClient.Close()

	slog.Info("Client created")

	// Auth
	//TODO: add proper credentials as secrets
	auth := smtp.PlainAuth(
		"",
		"", // login
		"", // pass
		"", // host
	)
	if err := smtpClient.Auth(auth); err != nil {
		slog.Error("Failed to authenticate", "error", err)
		return err
	}

	slog.Info("Authenticated")

	// Set the sender and recipient
	if err := smtpClient.Mail("foo@bar.com"); err != nil {
		slog.Error("Failed to set sender", "error", err)
		return err
	}

	slog.Info("Sender set")

	if err := smtpClient.Rcpt("bazz@fazz.com"); err != nil {
		slog.Error("Failed to set recipient", "error", err)
		return err
	}

	slog.Info("Recipient set")

	// Send the email body
	writer, err := smtpClient.Data()
	if err != nil {
		slog.Error("Failed to create writer", "error", err)
		return err
	}

	_, err = writer.Write([]byte(`TO: bazz@fazz.com \r\n
	Subject: Test email \r\n
	Testowy email \r\n
	`))
	if err != nil {
		slog.Error("Failed to write", "error", err)
		return err
	}

	err = writer.Close()
	if err != nil {
		slog.Error("Failed to close writer", "error", err)
		return err
	}

	return nil
}
