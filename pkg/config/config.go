// package config loads the enviroment variables for the app
package config

import (
	"fmt"
	"log"

	// autoload .env file if present
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// struct with config variables
type Config struct {
	DATABASE_URL string
	DB_TYPE      string
	DB_NAME      string
	DB_USER      string
	DB_PASSWORD  string
	DB_HOST      string
	DB_PORT      string
	REDIS_URL    string
	SECRET_KEY   string
	PORT         string
}

// load the env variables to the config struct
// returns error and config
// In Docker, environment variables are passed by the container runtime,
// so your Go app does not need godotenv inside the container.
func LoadConfigs() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error loading env variables: %v\n", err.Error())
	}
	viper.AutomaticEnv()

	config := &Config{
		DATABASE_URL: viper.GetString("DATABASE_URL"),
		DB_TYPE:      viper.GetString("DB_TYPE"),
		DB_NAME:      viper.GetString("DB_NAME"),
		DB_USER:      viper.GetString("DB_USER"),
		DB_PASSWORD:  viper.GetString("DB_PASSWORD"),
		DB_HOST:      viper.GetString("DB_HOST"),
		DB_PORT:      viper.GetString("DB_PORT"),
		REDIS_URL:    viper.GetString("REDIS_URL"),
		SECRET_KEY:   viper.GetString("SECRET_KEY"),
		PORT:         viper.GetString("PORT"),
	}

	// critical DB fields, causes panic if missing
	if config.DB_HOST == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}
	if config.DB_USER == "" || config.DB_PASSWORD == "" || config.DB_NAME == "" {
		return nil, fmt.Errorf("DB_USER, DB_PASSWORD, and DB_NAME are required")
	}

	return config, nil
}
