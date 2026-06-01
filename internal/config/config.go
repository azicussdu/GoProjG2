package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
}

func Load() (*Config, error) {
	err := godotenv.Load() // .env
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:       getEnv("PORT", "3030"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "azaaza"),
		DBPassword: getEnv("DB_PASSWORD", "123123"),
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
