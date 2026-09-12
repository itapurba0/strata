package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

func init() {
    err := godotenv.Load("../../.env")
    if err != nil {
        panic(err)
    }
}

func Load() (Config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return Config{}, fmt.Errorf("JWT_SECRET is not set")
	}

	return Config{
		DatabaseURL: databaseURL,
		Port:        port,
		JWTSecret:   secretKey,
	}, nil
}
