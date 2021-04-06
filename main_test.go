package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	// get vars for db init
	appdbUserName, appdbUserPassword, appdbName := getEnvVars()
	// Init app
	app.Initialize(
		appdbUserName,
		appdbUserPassword,
		appdbName)

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

// Testing the response to the /products endpoint with an empty table
func TestEmptyTable(t *testing.T) {
	// Delete all records from the products
	clearTable()

	// Execute the request
	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	// Testing the HTTP response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// Checking the body of the response and test that it is the textual representation of an empty array
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

// This function executes the request using the application’s router and returns the response
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

// This function testing the HTTP response code
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
