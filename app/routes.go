package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muanya/go-noter/db"
	"github.com/Muanya/go-noter/users"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Setup()

	defer db.Close()

	router := gin.Default()

	unauthorized := router.Group("")
	{
		unauthorized.GET("/health", users.Health)

	}

	usr := router.Group("/users")
	// usr.Use(auth.AuthMiddleWare())
	{
		usr.GET("/", users.GetAll)
		usr.POST("/", users.CreateUser)
		usr.GET("/:id", users.GetSingle)
	}

	// run server on a different goroutine
	go func() {
		if err := router.Run(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// block current goroutine, while listening for quit signal
	<-quit

	log.Println("Shutdown signal received, exiting...")
}
