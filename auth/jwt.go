package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// private
var (
	tokenValidHrs = 1 // token expires in 3 hrs
	host          = "127.0.0.1"
	secure        = false
	httpOnly      = false
	path          = "/"
)

// public
var (
	SigningMethod        *jwt.SigningMethodHMAC = jwt.SigningMethodHS256
	SecretKey                                   = "D#z+p@/9-$crt*k&y!"
	AuthorizationKeyword                        = "Authorization"
	BearerKeyword                               = "Bearer"
)

func Header(ctx *gin.Context) (string, error) {

	bearerToken := ctx.GetHeader(AuthorizationKeyword)

	if bearerToken == "" {
		return "", fmt.Errorf("missing authorization")
	}

	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != BearerKeyword {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

func GenerateToken(username string) (string, error) {
	mappedClaims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenValidHrs)).Unix(),
	}
	claims := jwt.NewWithClaims(SigningMethod, mappedClaims)

	return claims.SignedString([]byte(SecretKey))

}

func ParseClaim(c *gin.Context) (*jwt.MapClaims, error) {

	// Retrieve JWT token from cookie
	tkn, err := Header(c)

	if err != nil {
		return nil, err
	}

	// Parse JWT token with claims
	token, err := jwt.ParseWithClaims(tkn, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
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
	c.SetCookie(AuthorizationKeyword, "", -3600, path, host, secure, httpOnly)

}
