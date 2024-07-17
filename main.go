package main

import (
	"fmt"

	"github.com/Muanya/go-noter/users"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")
	setRoutes()
}

func setRoutes() {
	r := gin.Default()
	r.GET("/health", users.Health)
	r.Run()
}
