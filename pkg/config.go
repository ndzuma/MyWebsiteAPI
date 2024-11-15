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
	PORT        string
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func getAPIUsername() string {
	username := os.Getenv("API_USERNAME")
	if username == "" {
		username = "admin"
	}
	return username
}

func getAPIPassword() string {
	password := os.Getenv("API_PASSWORD")
	if password == "" {
		password = "password"
	}
	return password
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_PUBLIC_URL"),
		APIUsername: getAPIUsername(),
		APIPassword: getAPIPassword(),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		PORT:        getPort(),
	}, nil
}
