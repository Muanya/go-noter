package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	setup()
}

func setup() {
	conn, err := sql.Open("sqlite3", "./sqlite/note.db")

	if err != nil {
		fmt.Println("Error during conn opening")
		log.Fatal(err)
	}

	defer conn.Close()

	rows, err := conn.Query("select id, username, email from user")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("Current Users: ")

	for rows.Next() {
		var id int
		var uName, email string
		rows.Scan(&id, &uName, &email)
		fmt.Printf("%d: %s %s\n", id, uName, email)

	}
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK!",
	})
}
