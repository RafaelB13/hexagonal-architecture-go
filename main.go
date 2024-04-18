package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/rafaelb13/full-cycle-hexagonal/adapters/db"
	"github.com/rafaelb13/full-cycle-hexagonal/application"
	"log"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite3", "sqlite.db")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("Product 1", 10.2)
	if err != nil {
		log.Fatal(err)
	}

	productService.Enable(product)

}
