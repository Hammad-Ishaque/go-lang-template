package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:codexnow@tcp(127.0.0.1:3306)/Tutorials")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(err)
	con = db
	return db
}
