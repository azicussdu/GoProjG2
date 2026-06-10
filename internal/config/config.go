package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	err := godotenv.Load() // .env
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: getEnv("PORT", "3030"),
		Database: &DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "user123"),
			Password: getEnv("DB_PASSWORD", "qwerty123"),
			DBName:   getEnv("DB_NAME", "dbname"),
			SSLMode:  getEnv("SSL_MODE", "disable"),
		},
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
