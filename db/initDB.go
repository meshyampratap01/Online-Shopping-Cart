package db

import (
	"database/sql"
	"log"

	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
	    id TEXT PRIMARY KEY,
	    code TEXT NOT NULL UNIQUE,
	    discount REAL NOT NULL CHECK (discount > 0),
	    expiration_date TEXT NOT NULL
	);
	`

	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}	
}