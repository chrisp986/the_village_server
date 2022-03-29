package main

import (
	"log"

	"github.com/chrisp986/the_village_server/internal/database"
	"github.com/chrisp986/the_village_server/internal/models"
)

type application struct {
	players interface {
		Insert(models.Player) (int, error)
		Get(int) (*models.Player, error)
	}
}

func main() {
	log.Println("Starting the village server v0.1")

	err := initDB()
	if err != nil {
		log.Fatalln("Error initializing database:", err)
		return
	}

	db, err := connectDB()
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	initTables(db)

	app := &application{players: &database.PlayerModel{DB: db}}

	app.runServer()

}
