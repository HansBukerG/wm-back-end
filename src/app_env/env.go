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

	env_variable := os.Getenv(variable)
	if env_variable == "" {
		log.Fatal("Should define .env " + variable)
	}
	return env_variable
}
