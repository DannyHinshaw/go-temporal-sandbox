package temporal

import "os"

// GetHostPortEnv handles getting the correct temporal server url depending on context.
func GetHostPortEnv() string {
	url := os.Getenv("TEMPORAL_URL")
	if url == "" {
		url = "localhost:7233"
	}

	return url
}
