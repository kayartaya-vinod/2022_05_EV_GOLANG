package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createTable() {
	fmt.Println("Create table in sqlite3 db file")
	db, err := sql.Open("sqlite3", "trainingdb.sqlite")
	checkError(err)
	defer db.Close()

	sqlCmd := `create table shippers(
		shipper_id int primary key,
		company_name varchar(50) not null,
		phone varchar(50) unique
	)`

	db.Exec(sqlCmd)
	checkError(err)

} // db.Close() called here

func insertData() {
	db, err := sql.Open("sqlite3", "trainingdb.sqlite")
	checkError(err)
	defer db.Close()

	insertSql := "insert into shippers values(?, ?, ?)"
	id := 10
	name := "Speedy Express"
	phone := "5559201234"

	_, err = db.Exec(insertSql, id, name, phone)
	checkError(err)

	fmt.Println("Data inserted successfully")
}

func main() {
	insertData()
}
