package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetValue(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
