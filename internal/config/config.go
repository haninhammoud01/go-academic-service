package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret  string
	Expired time.Duration
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Parse JWT expiration
	jwtExpired, err := time.ParseDuration(getEnv("JWT_EXPIRED", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRED format: %w", err)
	}

	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "go-academic-service"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "academic_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:  getEnv("JWT_SECRET", "secret"),
			Expired: jwtExpired,
		},
	}, nil
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
