package main

import (
	"fmt"

	"github.com/chrisp986/the_village_server/internal/db"
	"github.com/chrisp986/the_village_server/internal/server"
)

func main() {
	fmt.Println("Starting the village server v0.1")
	db.InitDB()

	server.Run()

}
