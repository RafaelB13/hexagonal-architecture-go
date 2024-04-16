package db_test

import (
	"database/sql"
	"github.com/rafaelb13/full-cycle-hexagonal/adapters/db"
	"github.com/rafaelb13/full-cycle-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (id STRING, name STRING, price FLOAT, status STRING);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("abc", "Product 1", 10.2, "disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDb(DB)

	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.2, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDb(DB)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.4

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
