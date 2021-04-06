package main

import (
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

	appdbUserName, appdbUserPassword, appdbName := getEnvVars()

	app.Initialize(
		os.Getenv(appdbUserName),
		os.Getenv(appdbUserPassword),
		os.Getenv(appdbName))

	app.Run(":8889")
}
