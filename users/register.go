package users

import (
	"log"
	"net/http"

	"github.com/Muanya/go-noter/auth"
	"github.com/Muanya/go-noter/db"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Username  string `json:"user_name" binding:"required"`
	Firstname string `json:"first_name" binding:"required"`
	Lastname  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func RegisterUser(ctx *gin.Context) {

	var data RegisterRequest

	// Call BindJSON to bind the received JSON to newUser.
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newUser User

	err := newUser.GetFromRequest(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := GetHashPassword(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(newUser)

	// Add the new user to the database.
	result, err := db.Conn.Exec(
		"INSERT INTO user (username, email, firstname, lastname, password) VALUES (?, ?, ?, ?, ?)",
		newUser.Username, newUser.Email, newUser.Firstname, newUser.Lastname, string(password),
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

	// set the id of the new user
	newUser.Id = int(userId)

	// set jwt token in cookie
	if err = auth.SetCookieHandler(ctx, newUser.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, newUser)

}
