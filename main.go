package main

import (
	"fmt"

	"github.com/Muanya/go-noter/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")
	auth.Run()
}

func setRoutes() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
