package config

import (
	"os"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database *DBConfig
	JWT      *JWTConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func Load() (*Config, error) {
	err := godotenv.Load() // .env
	if err != nil {
		return nil, err
	}

	accessTTL, err := parseDurationEnv("JWT_ACCESS_TTL", "15m")
	if err != nil {
		return nil, err
	}

	refreshTTL, err := parseDurationEnv("JWT_REFRESH_TTL", "24h")
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
		JWT: &JWTConfig{
			SecretKey:  getEnv("JWT_SECRET", "secret123"),
			AccessTTL:  accessTTL,
			RefreshTTL: refreshTTL,
		},
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func parseDurationEnv(key, defaultValue string) (time.Duration, error) {
	value := getEnv(key, defaultValue) // value = "15m"
	dur, err := time.ParseDuration(value)
	if err != nil {
		return 0, apperrors.Internal("could not convert access token time to duration", err)
	}
	return dur, nil
}
