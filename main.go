package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Load environments
	godotenv.Load()

	if os.Getenv("PORT") == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Main program starting point
	fmt.Println("Starting application in PORT", os.Getenv("PORT"))
}
