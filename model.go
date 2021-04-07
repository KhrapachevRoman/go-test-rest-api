package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Struct to represent the `product`
type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Functions that deal with a single product as methods on struct `product`

// Get a product from the database by id
func (p *product) getProduct(db *sql.DB) error {
	//return errors.New("Not implemented")
	row := db.QueryRow("SELECT name, price FROM products WHERE id=?", p.ID)
	return row.Scan(&p.Name, &p.Price)
}

// Update a product in the database by id
func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=?, price=? WHERE id=?", p.Name, p.Price, p.ID)
	return err
}

// Delete a product in the database by id
func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=?", p.ID)
	return err
}

// Create a product in the database
func (p *product) createProduct(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO products (`name`, `price`) VALUES (?, ?)", p.Name, p.Price)

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(lastID)
	return nil
}

// Standalone function that fetches a list of products
func getProducts(db *sql.DB, start, count int) ([]product, error) {

	rows, err := db.Query("SELECT id, name, price FROM products LIMIT ? OFFSET ?", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
