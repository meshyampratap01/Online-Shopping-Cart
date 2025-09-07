package db

import (
	"database/sql"
	"log"

	_"github.com/mattn/go-sqlite3"
)


func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./shopping_cart.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatal("Error enabling foreign keys:", err)
	}

	createTables(db)
	seed(db)

	return db
}


func createTables(db *sql.DB) {
	createTables:=`
	CREATE TABLE IF NOT EXISTS users (
	    id TEXT PRIMARY KEY,
	    name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL,
	    role INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS products (
	    id TEXT PRIMARY KEY,
	    name TEXT NOT NULL,
	    price REAL NOT NULL CHECK (price >= 0),
	    stock INTEGER NOT NULL CHECK (stock >= 0)
	);

	CREATE TABLE IF NOT EXISTS cart (
	    id TEXT PRIMARY KEY,
	    user_id TEXT NOT NULL UNIQUE,
	    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS cart_items (
	    id TEXT PRIMARY KEY,
	    cart_id TEXT NOT NULL,
	    product_id TEXT NOT NULL,
	    quantity INTEGER NOT NULL CHECK (quantity > 0),
	    FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE,
	    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS coupons (
	    code TEXT NOT NULL UNIQUE,
	    discount REAL NOT NULL CHECK (discount > 0)
	);
	`

	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}	
}

func seed(db *sql.DB) {
	_, err := db.Exec(`
		INSERT OR IGNORE INTO users (id, name, email, password, role)
		VALUES ('u_Admin', 'Admin User', 'admin@shyam.com', 'admin@123', 1)
	`)
	if err != nil {
		log.Fatal("Error seeding admin:", err)
	}

	products := []struct {
		id    string
		name  string
		price float64
		stock int
	}{
		{"p1", "Laptop", 75000.00, 10},
		{"p2", "Smartphone", 35000.00, 25},
		{"p3", "Headphones", 2500.00, 50},
		{"p4", "Keyboard", 1200.00, 30},
		{"p5", "Monitor", 15000.00, 15},
	}

	for _, p := range products {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO products (id, name, price, stock)
			VALUES (?, ?, ?, ?)
		`, p.id, p.name, p.price, p.stock)
		if err != nil {
			log.Fatal("Error seeding products:", err)
		}
	}
}