package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app := App{}

	// Get the APP_DB_USERNAME environment variable
	appdbUserName, exists := os.LookupEnv("APP_DB_USERNAME")

	if exists {
		fmt.Println(appdbUserName)
	}

	// Get the APP_DB_PASSWORD environment variable
	appdbUserPassword, exists := os.LookupEnv("APP_DB_PASSWORD")

	if exists {
		fmt.Println(appdbUserPassword)
	}

	// Get the APP_DB_NAME environment variable
	appdbName, exists := os.LookupEnv("APP_DB_NAME")

	if exists {
		fmt.Println(appdbName)
	}

	app.Initialize(
		os.Getenv(appdbUserName),
		os.Getenv(appdbUserPassword),
		os.Getenv(appdbName))

	app.Run(":8889")
}
