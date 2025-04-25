// package config loads the enviroment variables for the app
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
}

// load the env variables to the config struct
// returns error and config
func LoadConfigs() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error initializing config: %v", err)
	}
	config := &Config{
		DATABASE_URL: getEnv("DATABASE_URL"),
		DB_TYPE:      getEnv("DB_TYPE"),
		DB_NAME:      getEnv("DB_NAME"),
		DB_USER:      getEnv("DB_USER"),
		DB_PASSWORD:  getEnv("DB_PASSWORD"),
		DB_HOST:      getEnv("DB_HOST"),
		DB_PORT:      getEnv("DB_PORT"),
		REDIS_URL:    getEnv("REDIS_URL"),
		SECRET_KEY:   getEnv("SECRET_KEY"),
	}
	// no error
	return config, nil
}

// check if .env has the correct keys
func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		log.Printf("missing %v key in .env\n", key)
	}
	return value
}
