package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}

}

func run(ctx context.Context) error {
	return nil
}
