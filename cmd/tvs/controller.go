package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chrisp986/the_village_server/internal/models"
	"github.com/gin-gonic/gin"
)

// func getPlayers(c *gin.Context) {
// 	if len(players) == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No players found!"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, players)
// }

func (a *application) getPlayer(c *gin.Context) {

	playerID := c.Param("player_id")
	pID, err := strconv.Atoi(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := a.players.Get(uint32(pID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

//Starts the flow to create a new player
func (a *application) createNewVillage(player_id uint32, player_name string) (uint32, error) {

	nv := models.Village{
		PlayerID:      player_id,
		VillageName:   fmt.Sprintf("Village %s", player_name),
		VillageSize:   100,
		VillageStatus: 0,
		VillageLocY:   uint32(randInt(0, 100)),
		VillageLocX:   uint32(randInt(0, 100)),
	}

	village_id, err := a.villages.Insert(nv)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Printf("New village created for: %s with village_id: %d", player_name, village_id)
	return village_id, err

}

// postPlayer creates a new player
func (a *application) postPlayer(c *gin.Context) {
	var newPlayer models.Player

	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new player to the players table
	player_id, err := a.players.Insert(newPlayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "players.Insert"})
		return
	}

	// Create a new village for the player
	village_id, err := a.createNewVillage(player_id, newPlayer.PlayerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "createNewVillage"})
		return
	}

	// Create a new village setup for the player
	_, err = a.villageSetup.InsertWithIDCheck(village_id, player_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "villageSetup.InsertWithIDCheck"})
		return
	}

	// Add the new player to the players_resources table
	village_id, err = a.villageResources.Insert(models.VillageResource{
		VillageID: village_id,
		PlayerID:  player_id,
		Food:      100,
		Wood:      100,
		Stone:     100,
		Metal:     100,
		Water:     100,
		Gold:      20,
	})

	c.JSON(http.StatusCreated, gin.H{
		"player_id":  player_id,
		"village_id": village_id,
	})
}

func (a *application) postConstructNewBuilding(c *gin.Context) {
	// this function is used to add a new building to a village
	// Get the village_id and the player_id as json
	// Add the new building to the building queue
	// When the construction is complete, add the building to the village
	//TODO: do this!

	log.Println("postConstructNewBuilding")

	// "village_id": 10,
	// "player_id": 1000,
	// "building_id": "h1",
	// "amount": 1,
	// "status": 0,
	// "start_time": "2020-01-01 00:00:00",
	// "finish_time": "2020-01-01 00:00:00"

	var newBuilding models.BuildingQueue

	if err := c.BindJSON(&newBuilding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "postConstructNewBuilding"})
		return
	}

	switch newBuilding.Status {
	case 1:
		log.Println("Start the construction of a new building")
		c.JSON(http.StatusOK, gin.H{"status": "Start the construction of a new building"})
		return
	case 2:
		log.Println("Start upgrading a building")
		c.JSON(http.StatusOK, gin.H{"status": "Start upgrading a building"})
		return
	default:
		log.Println("Error: Wrong status for building; 1 - start construction of a new building, 2 - upgrade existing building")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: Wrong status for building; 1 - start construction of a new building, 2 - upgrade existing building"})
		return
	}

}
