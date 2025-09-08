package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/meshyampratap01/OnlineShoppingCart/db"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/app"
)

func main() {
	db := db.InitDB()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	go func() {
		<-ch
		db.Close()
		fmt.Println("\nShutting down the server...")
		os.Exit(1)
	}()

	app := app.NewApp(db)

	app.Run()
}
