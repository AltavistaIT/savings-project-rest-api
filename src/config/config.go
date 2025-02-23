package config

import (
	"os"
	"strconv"
)

type serviceConfig struct {
	EnvCode int
	Port    string
}

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func ServiceConfig() *serviceConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	envCode, err := strconv.Atoi(os.Getenv("ENV_CODE"))
	if err != nil {
		envCode = 1
	}

	return &serviceConfig{
		EnvCode: envCode,
		Port:    port,
	}
}

func DatabaseConfig() *databaseConfig {
	values := &databaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return values
}
