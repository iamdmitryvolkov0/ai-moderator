package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	OpenAIAPIKey string
	ServerPort   string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load("app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppConfig = Config{
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s not set and no default provided", key)
		}
		return defaultValue
	}
	return value
}
