package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBName        string
	DBPort        string
	ServerPort    string
	MigrationsDir string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	return &Config{
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBHost:        os.Getenv("DB_HOST"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		MigrationsDir: os.Getenv("MIGRATIONS_DIR"),
	}
}
