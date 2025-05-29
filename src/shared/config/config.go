package config

import (
	"reflect"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	// Server
	EnvCode int    `mapstructure:"ENV_CODE"`
	Port    string `mapstructure:"PORT"`

	// Security
	SecurityAllowedOrigins []string `mapstructure:"SECURITY_ALLOWED_ORIGINS"`

	// Database
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.AutomaticEnv()

		t := reflect.TypeOf(Config{})
		for i := 0; i < t.NumField(); i++ {
			viper.BindEnv(t.Field(i).Tag.Get("mapstructure"), t.Field(i).Name)
		}

		config = &Config{}
		err := viper.Unmarshal(config)
		if err != nil {
			panic(err)
		}
	})
	return config
}
