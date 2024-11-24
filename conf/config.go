package conf

import (
	"os"
)

type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Config struct {
	EmailConfig    EmailConfig
	RabbitMQConfig RabbitMQConfig
}

func LoadConfig() *Config {
	return &Config{
		EmailConfig: EmailConfig{
			Host:     os.Getenv("EMAIL_HOST"),
			Port:     os.Getenv("EMAIL_PORT"),
			Username: os.Getenv("EMAIL_USERNAME"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
		RabbitMQConfig: RabbitMQConfig{
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
		},
	}
}
