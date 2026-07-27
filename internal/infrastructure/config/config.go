package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GeminiAPIKey string
	OpenAPIKey   string
	DefaultModel string
}

func readEnv(key string) string {
	return os.Getenv(key)
}

func Connection() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed connect env: %v", err)
	}

	return Config{
		GeminiAPIKey: readEnv("GEMINIAPIKEY"),
		OpenAPIKey:   readEnv("OPENAPIKEY"),
		DefaultModel: readEnv("DEFAULTMODEL"),
	}

}
