package configuration

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

/* Get an environment variable as a string!
 * If the environment variable is not set, the given default value is used!
 */
func getenvStr(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}

	return v
}

/* Get an environment variable as an integer!
 * If the environment variable is not set, the given default value is used!
 */
func getenvInt(key string, defaultValue int) int {
	s := getenvStr(key, "")
	v, err := strconv.Atoi(s)

	if err != nil {
		return defaultValue
	}

	return v
}

type Config struct {
	Port                       int
	Channel                    string // could be teams or mail
	NotificationFilterSeverity string // could be debug, info, warning, error
	LogLevel                   string // could be debug, info, warning, error
}

/* Load the configuration from the environment!
 * If a value is not set, the default value is used!
 */
func LoadConfig() (Config, error) {
	return Config{
		Port:                       getenvInt("PORT", 3000), //nolint:gomnd // port is 3000 by default
		Channel:                    getenvStr("CHANNEL", "teams"),
		NotificationFilterSeverity: getenvStr("SEVERITY", "warning"),
		LogLevel:                   getenvStr("LOG_LEVEL", "debug"),
	}, nil
}

func GetSlogLevel(logLevel string) slog.Level {
	if strings.EqualFold(logLevel, "info") {
		return slog.LevelInfo
	}

	if strings.EqualFold(logLevel, "warning") {
		return slog.LevelWarn
	}

	if strings.EqualFold(logLevel, "error") {
		return slog.LevelError
	}

	return slog.LevelDebug
}
