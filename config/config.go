package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser  	string
	DBPassword  string
	DBName 	    string
	DBHost 	    string
	DBPort 	    string
}


func NewConfig() *Config {
	return &Config{
		DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "duidQu"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
	}
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
