package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBPath         string
	JWTSecret      string
	JWTExpiryHours int
	AllowedOrigins string
}

func Load() *Config {
	var projectRoot string
	dir, err := os.Getwd()
	if err == nil {
		for {
			envPath := filepath.Join(dir, ".env")
			if _, err := os.Stat(envPath); err == nil {
				_ = godotenv.Load(envPath)
				projectRoot = dir
				break
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
	}

	if projectRoot == "" {
		projectRoot, _ = os.Getwd()
	}

	dbPath := getEnv("DB_PATH", "./database/ecommerce.db")
	if !filepath.IsAbs(dbPath) {
		dbPath = filepath.Join(projectRoot, dbPath)
	}

	expiryHours, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		expiryHours = 24
	}

	return &Config{
		Port:           getEnv("PORT", "8080"),
		DBPath:         dbPath,
		JWTSecret:      getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiryHours: expiryHours,
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5173"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
