package main

import (
	"fmt"
	"log"

	"github.com/chrisp986/the_village_server/internal/database"
	"github.com/chrisp986/the_village_server/internal/models"
)

type app struct {
	players interface {
		Insert(models.Player) (int, error)
	}
}

func main() {
	fmt.Println("Starting the village server v0.1")

	err := InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	db, err := ConnectDB()
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}

	fmt.Println("Connected to database")
	_ = &app{players: &database.PlayerModel{DB: db}}

	// var newPlayer = server.Player{
	// 	PlayerID:    6,
	// 	PlayerName:  "",
	// 	PlayerScore: 0,
	// 	Active:      false,
	// 	Connected:   false,
	// }

	// app.players.Insert(newPlayer)

	runServer()

}
