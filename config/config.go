package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AgentsBaseURL      string
	AgentsAPIKey       string
	AgentsDefaultModel string
}

var Settings *Config

func init() {
	Settings = Load()
}

func Load() *Config {
	_ = godotenv.Load()

	settings := &Config{
		AgentsBaseURL:      "http://localhost:8000",
		AgentsAPIKey:       "not-needed",
		AgentsDefaultModel: "gpt-oss-120b",
	}

	if v, ok := os.LookupEnv("AGENTS_BASE_URL"); ok {
		settings.AgentsBaseURL = v
	}

	if v, ok := os.LookupEnv("AGENTS_API_KEY"); ok {
		settings.AgentsAPIKey = v
	}

	if v, ok := os.LookupEnv("AGENTS_DEFAULT_MODEL"); ok {
		settings.AgentsDefaultModel = v
	}

	return settings
}
