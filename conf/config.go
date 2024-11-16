package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	EmailHost     string
	EmailPort     string
	EmailUsername string
	EmailPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		EmailHost:     os.Getenv("EMAIL_HOST"),
		EmailPort:     os.Getenv("EMAIL_PORT"),
		EmailUsername: os.Getenv("EMAIL_USERNAME"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
	}
}
