package main

import (
	"database/sql"

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
func (a *App) Initialize(user, password, dbname string) {}

// Simply start the application
func (a *App) Run(addr string) {}
