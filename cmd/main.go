package main

import (
	"github.com/meshyampratap01/OnlineShoppingCart/db"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/app"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	app := app.NewApp(db)

	app.Run()
}
