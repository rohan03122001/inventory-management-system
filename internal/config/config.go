package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	JWT_SECRET  string
}

func LoadConfig()(*Config, error) {
	// Load the .env file
	if err := godotenv.Load(); err !=nil{
		return nil , fmt.Errorf("Error Loading Environment File %w",err) 
	}	

	config := &Config{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	if config.DB_HOST == "" || config.DB_PORT == "" || config.DB_USER == "" || config.DB_PASSWORD == "" || config.DB_NAME == "" {
		return nil, fmt.Errorf("Error Database Config Missing")
	}

	return config, nil
}