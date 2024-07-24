package users

import (
	"net/http"

	"github.com/Muanya/go-noter/auth"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK!",
	})
}

func GetAll(ctx *gin.Context) {

	usrs, err := GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, usrs)

}

func GetUser(ctx *gin.Context) {

	// Get User from token

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
