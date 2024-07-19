package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func Setup() {
	var err error
	Conn, err = sql.Open("sqlite3", "./sqlite/note.db")

	if err != nil {
		fmt.Println("Error during conn opening")
		log.Fatal(err)
		panic(err)

	}
}

func Close() {
	if Conn != nil {
		log.Println("Closing DB Connection")
		Conn.Close()
	}
}
