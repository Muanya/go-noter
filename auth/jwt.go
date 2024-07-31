package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// private
var (
	tokenValidHrs = 3 // token expires in 3 hrs
	host          = "127.0.0.1"
	secure        = false
	httpOnly      = true
	path          = "/"
)

// public
var (
	SigningMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256
	SecretKey                            = "my-secret-key"
	TokenKeyword                         = "token"
)

func GenerateToken(username string) (string, error) {
	mappedClaims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenValidHrs)).Unix(),
	}
	claims := jwt.NewWithClaims(SigningMethod, mappedClaims)

	return claims.SignedString([]byte(SecretKey))

}

// Handler function to set a JWT cookie
func SetCookieHandler(c *gin.Context, username string) (string, error) {
	// Generate the JWT token
	tokenString, err := GenerateToken(username)
	if err != nil {
		return "", err
	}

	// Set the cookie
	c.SetCookie(TokenKeyword, tokenString, 3600*tokenValidHrs, path, host, secure, httpOnly)
	return tokenString, nil
}

func ParseClaim(c *gin.Context) (*jwt.MapClaims, error) {

	// Retrieve JWT token from cookie
	cookie, err := c.Cookie(TokenKeyword)

	if err != nil {
		return nil, err
	}

	// Parse JWT token with claims
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, gin.Error{}
	}

	return claims, nil

}

func ClearTokenHandler(c *gin.Context) {

	// Clear JWT token by setting an empty value and expired time in the cookie
	c.SetCookie(TokenKeyword, "", -3600, path, host, secure, httpOnly)

}
