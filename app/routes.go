package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muanya/go-noter/auth"
	"github.com/Muanya/go-noter/db"
	"github.com/Muanya/go-noter/users"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Setup()

	defer db.Close()

	router := gin.Default()
	router.Use(auth.CORSMiddleware())

	unauthorized := router.Group("")
	{
		unauthorized.GET("/health", users.Health)
		unauthorized.POST("/login", users.LoginUser)
		unauthorized.POST("/register", users.RegisterUser)
		unauthorized.POST("/logout", users.Logout)

	}

	usr := router.Group("/users")
	usr.Use(auth.JWTVerifyMiddleWare())
	{
		usr.GET("/", users.GetUser)
		usr.GET("/all", users.GetAll)
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
