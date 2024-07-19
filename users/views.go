package users

import (
	"log"
	"net/http"

	"github.com/Muanya/go-noter/db"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK!",
	})
}

func GetAll(ctx *gin.Context) {
	rows, err := db.Conn.Query("SELECT id, username, email, firstname, lastname FROM user")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	usrs, err := FormatRowsToUsers(rows)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usrs)

}

func GetSingle(ctx *gin.Context) {

	userId := ctx.Param("id")

	rows, err := db.Conn.Query("SELECT id, username, email, firstname, lastname FROM user WHERE id = ?", userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usrs, err := FormatRowsToUsers(rows)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(usrs) > 1 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Multiple Users found"})
	} else {
		ctx.JSON(http.StatusOK, usrs[0])
	}

}

func CreateUser(ctx *gin.Context) {

	var data Response

	// Call BindJSON to bind the received JSON to newUser.
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := GetUsersFromRequest(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := GetUserPassword(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(newUser)

	// Add the new user to the database.
	result, err := db.Conn.Exec(
		"INSERT INTO user (username, email, firstname, lastname, password) VALUES (?, ?, ?, ?, ?)",
		newUser.Username, newUser.Email, newUser.Firstname, newUser.Lastname, password,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(result.LastInsertId())
	userId, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser.Id = int(userId)

	ctx.JSON(http.StatusCreated, newUser)

}
