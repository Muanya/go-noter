package users

import (
	"net/http"

	"github.com/Muanya/go-noter/auth"
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

func GetUser(ctx *gin.Context) {

	claims, err := auth.ParseClaim(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not Authorized!"})
		return
	}

	var user User

	// Extract user ID from claims
	username, _ := (*claims)["sub"].(string)

	// Query user from database using ID
	if err = user.GetByUsername(username); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not Authorized!"})
		return
	}

	// Return user details as JSON response
	ctx.JSON(http.StatusOK, user)
}

func Logout(ctx *gin.Context) {

	auth.ClearTokenHandler(ctx)

	// Return success response indicating logout was successful
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})

}
