package notifications

import (
	"context"

	"github.com/harmoniemand/alerting-demo/internal/configuration"
)

type Channel interface {
	SendMessage(ctx context.Context, n Notification) error
}

type NotificationManager struct {
	Config  configuration.Config
	Channel Channel
}

/* Create a new notification manager with the given configuration!
 * Adds the given channel to the manager!
 */
func NewNotificationManager(config configuration.Config, channel Channel) (NotificationManager, error) {
	return NotificationManager{
		Config:  config,
		Channel: channel,
	}, nil
}

/* Convert a notification level to a number!
 * The number can be used to easily compare the severity of two notifications!
 */
func (m *NotificationManager) LevelAsNumber(level string) int {
	levels := []string{"debug", "info", "warning", "error"}

	for i, n := range levels {
		if level == n {
			return i
		}
	}

	return len(levels)
}

/* Check if the severity of level1 is lower than the severity of level2!
 * This can be used to filter notifications by severity!
 */
func (m *NotificationManager) IsLowerThan(level1 string, level2 string) bool {
	return m.LevelAsNumber(level1) < m.LevelAsNumber(level2)
}

/* Send a notification via the notification channel!
 * If the notification is lower than the configured severity filter, the notification will not be sent!
 */
func (m *NotificationManager) SendNotification(ctx context.Context, notification Notification) error {
	if !m.IsLowerThan(notification.Type, m.Config.NotificationFilterSeverity) {
		return nil
	}

	return m.Channel.SendMessage(ctx, notification)
}
