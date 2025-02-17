package config

import (
	"os"
	"strconv"
)

type serviceConfig struct {
	EnvCode int
	Port    string
}

func Env() *serviceConfig {
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
