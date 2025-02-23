package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pzapasnik/email-sender/internal/handler/index"
	renderer "github.com/pzapasnik/email-sender/internal/renderer/templ"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	router := gin.New()
	router.Use(sloggin.New(logger))

	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// This is bad. You can't normally set your own renderer in gin
	// So you have to overwrite it by your self having in mind that
	// order of setting settings on router matters
	// It would be nice to have functional options pattern here
	router.HTMLRender = renderer.NewTemplRenderer(router.HTMLRender)

	router.GET("/html", index.New().Handle)
	router.GET("/templ", index.New().HandleTempl)

	//TODO: add custom server timeouts and other settings
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//TODO: add errgroup and graceful shutdown
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	return nil
}
