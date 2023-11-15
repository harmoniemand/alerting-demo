package configuration

import (
	"os"
	"strconv"
)

type Config struct {
	Port                       int
	Channel                    string // could be teams or mail
	NotificationFilterSeverity string // could be debug, info, warning, error
}

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

/* Load the configuration from the environment!
 * If a value is not set, the default value is used!
 */
func LoadConfig() (Config, error) {
	return Config{
		Port:                       getenvInt("PORT", 3000), //nolint:gomnd // port is 3000 by default
		Channel:                    getenvStr("CHANNEL", "teams"),
		NotificationFilterSeverity: getenvStr("SEVERITY", "warning"),
	}, nil
}
