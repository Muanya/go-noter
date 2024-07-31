package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader != "secret-token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			ctx.Abort()
			return

		}

		ctx.Next()
	}
}

func JWTVerifyMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader, err := ctx.Cookie(TokenKeyword)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization"})
			ctx.Abort()
			return

		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You're Unauthorized!"})
				ctx.Abort()

				return nil, fmt.Errorf("")

			}

			return []byte(SecretKey), nil

		})

		if err != nil {
			fmt.Println(err)

			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You're Unauthorized due to error parsing the JWT"})
			ctx.Abort()
			return

		}

		if token.Valid {
			ctx.Next()

		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You're Unauthorized due to invalid token"})

		}

	}
}
