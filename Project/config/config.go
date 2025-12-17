package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() Config {
	return Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
