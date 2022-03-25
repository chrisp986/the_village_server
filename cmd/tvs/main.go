package main

import (
	"fmt"

	"github.com/chrisp986/the_village_server/internal/db"
	"github.com/chrisp986/the_village_server/internal/server"
)

func main() {
	fmt.Println("Starting the village server v0.1")

	err := db.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	server.Run()

}
