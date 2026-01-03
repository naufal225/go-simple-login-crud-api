package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	return &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASS", "nma225"),
		DBName:    getEnv("DB_NAME", "go_simple_api"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
		JWTSecret: getEnv("JWT_SECRET", "indonesia-bubar-2030-presiden-lu-sehat-bang-?-presiden-gw-nanem-sawit-jir"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
