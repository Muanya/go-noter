package users

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Muanya/go-noter/auth"
	"github.com/Muanya/go-noter/db"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(ctx *gin.Context) {
	var request LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validPassword string
	err := db.Conn.QueryRow("SELECT password FROM user WHERE username = ?", request.Username).Scan(&validPassword)

	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else if err != nil {
		log.Println("Error querying database:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	isValid := CompareHashPassword(request.Password, validPassword)

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// set jwt token in cookie
	if err = auth.SetCookieHandler(ctx, request.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}
