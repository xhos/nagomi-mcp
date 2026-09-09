package config

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type Config struct {
	NagomiCoreURL string
	APIKey        string

	UserID        uuid.UUID
	ListenAddress string
	BaseURL       string

	LogFormat string // "json" | "text"
	LogLevel  log.Level
}

// parseAddress handles "8080", ":8080", or "127.0.0.1:8080".
func parseAddress(s string) string {
	s = strings.TrimSpace(s)
	if strings.Contains(s, ":") {
		return s
	}
	return ":" + s
}

func Load() Config {
	nagomiCoreURL := os.Getenv("NAGOMI_CORE_URL")
	if nagomiCoreURL == "" {
		panic("NAGOMI_CORE_URL environment variable is required")
	}

	apiKey := os.Getenv("NAGOMI_API_KEY")
	if apiKey == "" {
		panic("NAGOMI_API_KEY environment variable is required")
	}

	userID := os.Getenv("NAGOMI_MCP_USER_ID")
	if userID == "" {
		panic("NAGOMI_MCP_USER_ID environment variable is required")
	}

	baseURL := strings.TrimRight(os.Getenv("NAGOMI_MCP_BASE_URL"), "/")
	if baseURL == "" {
		panic("NAGOMI_MCP_BASE_URL environment variable is required")
	}

	listenAddress := os.Getenv("NAGOMI_MCP_LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "127.0.0.1:55553"
	}

	logLevel, err := log.ParseLevel(os.Getenv("NAGOMI_MCP_LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	logFormat := strings.ToLower(strings.TrimSpace(os.Getenv("NAGOMI_MCP_LOG_FORMAT")))
	if logFormat != "json" {
		logFormat = "text"
	}

	return Config{
		NagomiCoreURL: nagomiCoreURL,
		APIKey:        apiKey,
		UserID:        uuid.MustParse(userID),
		ListenAddress: parseAddress(listenAddress),
		BaseURL:       baseURL,
		LogLevel:      logLevel,
		LogFormat:     logFormat,
	}
}
