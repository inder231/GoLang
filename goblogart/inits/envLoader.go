package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file into the environment variables.
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}