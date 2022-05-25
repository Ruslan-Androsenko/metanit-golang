package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// addingData(db)
	// gettingData(db)
	// updateData(db)
	// deleteData(db)
}

func addingData(db *sql.DB) {
	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)", "Apple", 72000)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId()) // id не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	// добавление записи в таблицу, с возвращением id-добавленной записи
	var id int
	db.QueryRow("insert into Products (model, company, price) values ('Mate 10 Pro', $1, $2) returning id",
		"Huawei", 35000).Scan(&id)
	fmt.Println(id)
}

func gettingData(db *sql.DB) {
	rows, err := db.Query("select * from Products")

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

	row := db.QueryRow("select * from Products where id = $1", 2)
	prod := product{}
	err = row.Scan(&prod.id, &prod.model, &prod.company, &prod.price)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println(prod.id, prod.model, prod.company, prod.price)
}

func updateData(db *sql.DB) {
	// обновляем строку с id = 1
	result, err := db.Exec("update Products set price = $1 where id = $2", 69000, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected()) // количество обновленных строк
}

func deleteData(db *sql.DB) {
	result, err := db.Exec("delete from Products where id = $1", 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
