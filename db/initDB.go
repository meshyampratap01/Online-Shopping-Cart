package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully!")

	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL,
		cart JSON
	);`
	_, err := db.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}
	productTable := `
	CREATE TABLE IF NOT EXISTS products (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		stock INT NOT NULL
	);`
	_, err = db.Exec(productTable)
	if err != nil {
		log.Fatal(err)
	}

	couponTable := `
	CREATE TABLE IF NOT EXISTS coupons (
		code VARCHAR(50) PRIMARY KEY,
		discount DECIMAL(5, 2) NOT NULL,
	);`
	_, err = db.Exec(couponTable)
	if err != nil {
		log.Fatal(err)
	}
}
