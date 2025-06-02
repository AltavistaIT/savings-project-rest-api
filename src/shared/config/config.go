package config

import (
	"reflect"
	"sync"

	"github.com/spf13/viper"
	"github.com/ssssshel/sp-api/src/shared/logger"
)

type Config struct {
	// Server
	Port string `mapstructure:"PORT"`

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
		viper.SetConfigFile(".env")
		viper.AddConfigPath(".") // Si ejecutas desde raíz del proyecto
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal("❌ No se pudo leer el archivo .env: ", err)
		}
		viper.AutomaticEnv()

		t := reflect.TypeOf(Config{})
		for i := 0; i < t.NumField(); i++ {
			key := t.Field(i).Tag.Get("mapstructure")
			_ = viper.BindEnv(key)
			logger.Info("🔍 %s = %v", key, viper.Get(key)) // debug opcional
		}

		config = &Config{}
		err := viper.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		logger.Info("✅ Config cargada: %+v", config)
	})
	return config
}
