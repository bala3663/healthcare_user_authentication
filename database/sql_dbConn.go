package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	var err error

	db, err = sql.Open("mysql", "root:india@123@tcp(localhost:3306)/Patient_information")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected Successfully")

	//create table
}
