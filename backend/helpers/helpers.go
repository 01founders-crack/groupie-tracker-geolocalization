package helpers

import (
	"fmt"
	"os"
	"strings"
)

// Load environment variables from a file
func LoadEnv(envFile string) error {
	content, err := os.ReadFile(envFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}

	return nil
}

func InitEnv() (string, string) {
	// Load environment variables from .env file
	if err := LoadEnv(".env"); err != nil {
		fmt.Println("Error loading .env file:", err)

	}
	accessToken := os.Getenv("ACCESS_TOKEN")
	gMapsToken := os.Getenv("GMAPS_TOKEN")

	// Check if the access token is empty or not set
	if accessToken == "" || gMapsToken == "" {
		fmt.Println("Access token not found in environment variable ACCESS_TOKEN or GMAPS_TOKEN")
		// Handle the case where the access token is missing or empty
		return "", ""
	}
	return accessToken, gMapsToken
}
