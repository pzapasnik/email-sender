package email

import (
	"log/slog"
	"net/http"

	"github.com/pzapasnik/email-sender/web/templates"
)

func GetRootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t := templates.Root(false)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := t.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func PostSendHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		email := r.FormValue("email")

		slog.Info("parsing form", "r.Form", r.Form, "email", email)

		t := templates.Root(true)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := t.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
