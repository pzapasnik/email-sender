package stdout

import (
	"log/slog"

	"github.com/pzapasnik/email-sender/internal/service"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (e *Service) Send(mail service.MailDTO) error {
	slog.Info("Mail Send", "Mail", mail)
	return nil
}
