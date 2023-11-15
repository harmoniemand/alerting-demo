package main

import (
	"log/slog"
	"os"

	"github.com/harmoniemand/alerting-demo/internal/api"
	"github.com/harmoniemand/alerting-demo/internal/configuration"
)

func main() {
	opts := slog.HandlerOptions{
		Level: configuration.GetSlogLevel(os.Getenv("LOG_LEVEL")),
	}

	textHandler := slog.NewTextHandler(os.Stdout, &opts)
	logger := slog.New(textHandler)
	slog.SetDefault(logger)

	slog.Info("starting application alerting-demo")

	config, err := configuration.LoadConfig()
	if err != nil {
		slog.Error("Error loading config: %v", err)
		os.Exit(1)
	}

	server, err := api.NewServer(config)
	if err != nil {
		slog.Error("Error creating server: %v", err)
		os.Exit(1)
	}

	if err = server.Start(); err != nil {
		slog.Error("Error starting server: %v", err)
	}

	os.Exit(0)
}
