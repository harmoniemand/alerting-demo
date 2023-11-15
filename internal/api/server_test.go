package api_test

import (
	"testing"

	"github.com/harmoniemand/alerting-demo/internal/api"
	"github.com/harmoniemand/alerting-demo/internal/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateServer(t *testing.T) {
	config := configuration.Config{
		Port:    3000,
		Channel: "teams",
	}

	server, err := api.NewServer(config)

	require.NoError(t, err)
	assert.NotNil(t, server)
}
