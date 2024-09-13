package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
}

func LoadConfig(envFile string) (*Config, error) {
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			return nil ,err
		}
	}

	return &Config{
		AppPort: getEnv("APP_PORT", "8000"),		
	}, nil
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}