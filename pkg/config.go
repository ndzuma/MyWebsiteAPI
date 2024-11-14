package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseUrl string
	APIUsername string
	APIPassword string
	JwtSecret   string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_PUBLIC_URL"),
		APIUsername: os.Getenv("API_USERNAME"),
		APIPassword: os.Getenv("API_PASSWORD"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
