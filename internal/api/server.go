package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/harmoniemand/alerting-demo/internal/channels"
	"github.com/harmoniemand/alerting-demo/internal/configuration"
	"github.com/harmoniemand/alerting-demo/internal/notifications"
)

type Server struct {
	Config              configuration.Config
	NotificationHandler notifications.NotificationHandler
}

/* Create a new server with the given configuration!
 * Creates all dependencies using the given configuration!
 */
func NewServer(config configuration.Config) (*Server, error) {
	slog.Debug("Creating server")

	channel, err := channels.NewTeamsChannel(config)
	if err != nil {
		slog.Error("Error creating teams channel: %v", err)
		os.Exit(1)
	}

	manager, err := notifications.NewNotificationManager(config, channel)
	if err != nil {
		slog.Error("Error creating notification manager: %v", err)
		os.Exit(1)
	}

	notificationHandler, err := notifications.NewNotificationHandler(config, manager)
	if err != nil {
		slog.Error("Error on creating NotificationHandler: %v", err)
		return nil, err
	}

	return &Server{
		Config:              config,
		NotificationHandler: notificationHandler,
	}, nil
}

/* Start the server!
 * Creates a new router and adds the notification handler to it!
 * Starts the server on the given port!
 */
func (s *Server) Start() error {
	slog.Info("Starting server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/", s.NotificationHandler.HandlePostNotification)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%v", s.Config.Port),
		WriteTimeout:      1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Error on starting server: %v", err)
		return err
	}

	return nil
}
