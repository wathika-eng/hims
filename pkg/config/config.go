// package config loads the enviroment variables for the app
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// struct with config variables
type Config struct {
	DATABASE_URL string
	DB_TYPE      string
	DB_NAME      string
	DB_USER      string
	DB_PASS      string
	DB_HOST      string
	DB_PORT      string
	REDIS_URL    string
	SECRET_KEY   string
}

// load the env variables to the config struct
// returns error and config
func LoadConfigs() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error initializing config: %v", err)
	}
	config := &Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		DB_TYPE:      os.Getenv("DB_TYPE"),
		DB_NAME:      os.Getenv("DB_NAME"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASS:      os.Getenv("DB_PASS"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_PORT:      os.Getenv("DB_PORT"),
		REDIS_URL:    os.Getenv("REDIS_URL"),
		SECRET_KEY:   os.Getenv("SECRET_KEY"),
	}
	// no error
	return config, nil
}
