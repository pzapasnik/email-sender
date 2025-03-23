package mail

import (
	"net/http"

	"github.com/pzapasnik/email-sender/internal/service"
	"github.com/pzapasnik/email-sender/web/templates"
)

type Handler struct {
	service service.MailService
}

func NewHandler(mailService service.MailService) *Handler {
	return &Handler{service: mailService}
}

func (h *Handler) HandleGetEmailForm() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t := templates.Mail(false)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := t.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) HandlePostEmailSend() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		mail := service.MailDTO{
			Sender:     "Pawel.zapasnik@visionary-consulting.pl",
			Recipients: []string{r.PostFormValue("to")},
			Subject:    r.PostFormValue("subject"),
			Body:       r.PostFormValue("message"),
		}

		err = h.service.Send(mail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		t := templates.Mail(true)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := t.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
