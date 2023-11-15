package channels

import (
	"context"
	"log/slog"

	"github.com/harmoniemand/alerting-demo/internal/configuration"
	"github.com/harmoniemand/alerting-demo/internal/notifications"
)

type TeamsChannel struct {
	Config configuration.Config
}

func NewTeamsChannel(config configuration.Config) (TeamsChannel, error) {
	return TeamsChannel{
		Config: config,
	}, nil
}

/* Send a message to teams!
 * This is a stub implementation!
 */
func (c TeamsChannel) SendMessage(ctx context.Context, n notifications.Notification) error {
	slog.InfoContext(ctx, "sending notification via teams channel: ", "type", n.Type, "name", n.Name, "description", n.Description)

	// ToDO: implement sending notification to teams via webhooks or something similar
	// https://learn.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using?tabs=cURL

	return nil
}
