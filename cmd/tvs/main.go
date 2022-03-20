package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Player struct {
	PlayerID    string `json:"player_id"`
	PlayerName  string `json:"player_name"`
	PlayerScore int64  `json:"player_score"`
	Active      bool   `json:"active"`
}

var players = []Player{
	{PlayerID: "1", PlayerName: "John Doe", PlayerScore: 100, Active: true},
	{PlayerID: "2", PlayerName: "Jane Doe", PlayerScore: 200, Active: true},
}

func main() {
	fmt.Println("Starting the village server v0.1")
	runServer()
}

func runServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	router.GET("/players", getPlayers)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func getPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}
