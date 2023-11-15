package notifications

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/harmoniemand/alerting-demo/internal/configuration"
)

type NotificationHandler struct {
	Config  configuration.Config
	Manager NotificationManager
}

/* Create a new notification handler with the given configuration!
 * Adds the given notification manager to the handler!
 */
func NewNotificationHandler(config configuration.Config, manager NotificationManager) (NotificationHandler, error) {
	return NotificationHandler{
		Config:  config,
		Manager: manager,
	}, nil
}

/* Handle a post request to the notification endpoint!
 * Decodes the request body to a notification struct!
 * Sends the notification via the notification manager!
 */
func (h *NotificationHandler) HandlePostNotification(w http.ResponseWriter, r *http.Request) {
	slog.DebugContext(r.Context(), "Running HandlePostNotification")

	var notification Notification

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.ErrorContext(r.Context(), "Error on decoding json body: %v", err)

		if _, e := w.Write([]byte("Bad Request")); e != nil {
			slog.ErrorContext(r.Context(), "Error on writing request: %v", e)
			return
		}
	}

	err = h.Manager.SendNotification(r.Context(), notification)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error on sending notification via channel: %v", err)

		if _, e := w.Write([]byte("internal server error")); e != nil {
			slog.ErrorContext(r.Context(), "Error on writing request: %v", e)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
