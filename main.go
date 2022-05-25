package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	db, err := sql.Open("mysql", "root:root@/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// addingData(db)
	gettingData(db)
}

func addingData(db *sql.DB) {
	result, err := db.Exec("insert into productdb.products (model, company, price) values (?, ?, ?)", "iPhone X", "Apple", 72000)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func gettingData(db *sql.DB) {
	rows, err := db.Query("select * from productdb.products")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)

		if err != nil {
			fmt.Println(err)
			continue
		}

		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}

	row := db.QueryRow("select * from productdb.products where id = ?", 2)
	prod := product{}
	err = row.Scan(&prod.id, &prod.model, &prod.company, &prod.price)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println(prod.id, prod.model, prod.company, prod.price)
}
