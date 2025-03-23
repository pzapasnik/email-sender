package smtp

import (
	"crypto/tls"
	"log/slog"
	"net/smtp"

	"github.com/pzapasnik/email-sender/internal/service"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (e *Service) Send(mail service.MailDTO) error {
	//SMTP conn
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "server189430.nazwa.pl",
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
		"Pawel.zapasnik@visionary-consulting.pl", // login
		"Pawel@nazwa12",                          // pass
		"server189430.nazwa.pl",                  // host
	)
	if err := smtpClient.Auth(auth); err != nil {
		slog.Error("Failed to authenticate", "error", err)
		return err
	}

	slog.Info("Authenticated")

	// Set the sender and recipient
	if err := smtpClient.Mail("Pawel.zapasnik@visionary-consulting.pl"); err != nil {
		slog.Error("Failed to set sender", "error", err)
		return err
	}

	slog.Info("Sender set")

	if err := smtpClient.Rcpt("pzapasnik@gmail.com"); err != nil {
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
