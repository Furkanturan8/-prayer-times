package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct holds configuration values
type Config struct {
	MySQLUsername string
	MySQLPassword string
	MySQLHost     string
	MySQLPort     string
	MySQLDBName   string
	SecretKey     string
	EmailAddress  string
	EmailPassword string
	SMTPPort      int
	SMTPServer    string
	Port          string
}

// Load loads configuration from environment variables or .env file
func Load() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load("backend/.env")
	if err != nil {
		log.Println("Error loading .env file, using default environment variables.")
	}

	// Parse SMTP port as an integer
	smtpPort, err := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	if err != nil {
		log.Println("Invalid SMTP port, using default 587.")
		smtpPort = 587
	}

	// Initialize configuration struct
	config := &Config{
		MySQLUsername: getEnv("MYSQL_USERNAME", "username"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "password"),
		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLDBName:   getEnv("MYSQL_DBNAME", "dbname"),
		SecretKey:     getEnv("SECRET_KEY", "default_secret_key"),
		EmailAddress:  getEnv("EMAIL_ADDRESS", "default_email@example.com"),
		EmailPassword: getEnv("EMAIL_PASSWORD", "default_email_password"),
		SMTPPort:      smtpPort,
		SMTPServer:    getEnv("SMTP_SERVER", "smtp.example.com"),
		Port:          getEnv("PORT", "3000"),
	}

	return config, nil
}

// getEnv retrieves the value of an environment variable, with a fallback default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
