package main

import (
	"log"

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
	appdbUserName, appdbUserPassword, appdbName := getEnvVars()
	app := App{}
	app.Initialize(
		appdbUserName,
		appdbUserPassword,
		appdbName)

	app.Run(":8889")
}
