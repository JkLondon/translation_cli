package main

import (
	"log"

	"github.com/JkLondon/translation_cli/cmd"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cmd.Execute()
}
