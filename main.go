package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	addingData(db)
}

func addingData(db *sql.DB) {
	result, err := db.Exec("insert into productdb.products (model, company, price) values (?, ?, ?)", "iPhone X", "Apple", 72000)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
