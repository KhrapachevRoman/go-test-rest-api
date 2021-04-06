package main

import (
	"log"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	appdbUserName, appdbUserPassword, appdbName := getEnvVars()

	// Init app
	app.Initialize(
		os.Getenv(appdbUserName),
		os.Getenv(appdbUserPassword),
		os.Getenv(appdbName))

	// Make sure that the table we need for testing is available
	ensureTableExists()
	// Executing all the tests
	code := m.Run()
	// Clean the database
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM products")
	app.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
