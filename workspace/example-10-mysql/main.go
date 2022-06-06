package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Start of main()")

	db, err := sql.Open("mysql", "root:Welcome#123@tcp(localhost)/northwind")
	checkError(err)
	defer db.Close() // executes after the last statement in this function

	log.Println("There was no error connecting to mysql database server")
	rows, err := db.Query("select * from SHIPPERS")
	checkError(err)

	for rows.Next() {
		var id int
		var name, phone string
		rows.Scan(&id, &name, &phone)
		log.Printf("Id = %v, Name = %v, Phone = %v", id, name, phone)
	}

	log.Println("End of main()")
}
