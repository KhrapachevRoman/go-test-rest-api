package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// This struct exposes references to the router and the database that the application uses.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// The Initialize method will take in the details required to connect to the database.
// It will create a database connection and wire up the routes to respond according to the requirements.
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:8889)/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

// Simply start the application
func (a *App) Run(addr string) {}

// Get environment variables
func getEnvVars() (user, password, dbname string) {
	// Get the APP_DB_USERNAME environment variable
	user, exists := os.LookupEnv("APP_DB_USERNAME")

	if !exists {
		log.Fatal("No APP_DB_USERNAME in .env file found")
	}

	// Get the APP_DB_PASSWORD environment variable
	password, exists = os.LookupEnv("APP_DB_PASSWORD")

	if !exists {
		log.Fatal("No APP_DB_PASSWORD in .env file found")
	}

	// Get the APP_DB_NAME environment variable
	dbname, exists = os.LookupEnv("APP_DB_NAME")

	if !exists {
		log.Fatal("No APP_DB_NAME in .env file found")
	}

	return
}
