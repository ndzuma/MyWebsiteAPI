package config

import (
	"fmt"
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
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "dev"
	}

	var err error
	if mode == "prod" {
		err = godotenv.Load("ENV")
	} else {
		err = godotenv.Load(".env")
	}

	if err != nil && mode != "prod" {
		return nil, fmt.Errorf("error loading env file: %w", err)
	}

	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_PUBLIC_URL"),
		APIUsername: getAPIUsername(),
		APIPassword: getAPIPassword(),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		PORT:        getPort(),
	}, nil
}
