package config

type Config struct {
	AgentsBaseURL string    `default: "http://192.168.1.202:8080"`
	AgensAPIKey string		`default: "not-needed"`
	AgentsModel string		`default: "gpt-oss-120b"`
}

func Load() *Config {

	cfg := &Config{}

	if v :
}
