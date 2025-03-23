package service

type MailDTO struct {
	Sender     string
	Recipients []string
	Subject    string
	Body       string
}

type MailService interface {
	Send(mail MailDTO) error
}
