package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	mailhandler "github.com/pzapasnik/email-sender/internal/handler/mail"
	mailsmtpservice "github.com/pzapasnik/email-sender/internal/service/mail/smtp"
	mailstdoutservice "github.com/pzapasnik/email-sender/internal/service/mail/stdout"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context) error {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web"))

	_ = mailhandler.NewHandler(mailsmtpservice.NewService())
	mailHandler := mailhandler.NewHandler(mailstdoutservice.NewService())

	mux.Handle("/static/", fs)
	mux.HandleFunc("/mail", mailHandler.HandleGetEmailForm())
	mux.HandleFunc("/mail-send", mailHandler.HandlePostEmailSend())

	//TODO: add custom server timeouts and other settings
	server := &http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelDebug),
	}

	slog.Info("running server")
	//TODO: add errgroup and graceful shutdown
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	return nil
}
