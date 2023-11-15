package notifications_test

import (
	"testing"

	"github.com/harmoniemand/alerting-demo/internal/configuration"
	"github.com/harmoniemand/alerting-demo/internal/notifications"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/* checks, if all the levels are correctly mapped to numbers!
 */
func TestLevelAsNumber(t *testing.T) {
	config := configuration.Config{
		Port:                       3000,
		Channel:                    "teams",
		NotificationFilterSeverity: "warning",
	}
	manager, err := notifications.NewNotificationManager(config, nil)
	require.NoError(t, err)

	assert.Equal(t, 0, manager.LevelAsNumber("debug"))
	assert.Equal(t, 1, manager.LevelAsNumber("info"))
	assert.Equal(t, 2, manager.LevelAsNumber("warning"))
	assert.Equal(t, 3, manager.LevelAsNumber("error"))
	assert.Equal(t, 4, manager.LevelAsNumber("unknown"))
}

/* checks, if the severity filter is working correctly!
 */
func TestIsSeverityLowerThan(t *testing.T) {
	config := configuration.Config{
		Port:                       3000,
		Channel:                    "teams",
		NotificationFilterSeverity: "warning",
	}
	manager, err := notifications.NewNotificationManager(config, nil)
	require.NoError(t, err)

	assert.True(t, manager.IsLowerThan("debug", "info"))
	assert.True(t, manager.IsLowerThan("info", "warning"))
	assert.True(t, manager.IsLowerThan("warning", "error"))
	assert.False(t, manager.IsLowerThan("error", "warning"))
	assert.False(t, manager.IsLowerThan("warning", "info"))
	assert.False(t, manager.IsLowerThan("info", "debug"))
}
