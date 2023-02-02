package app_env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(variable string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVariable := os.Getenv(variable)
	if envVariable == "" {
		log.Fatal("Should define .env " + variable)
	}
	return envVariable
}
