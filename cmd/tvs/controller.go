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

	player, err := a.players.Get(pID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

//Starts the flow to create a new player
func (a *application) createNewVillage(player_id int32, player_name string) (int, error) {

	nv := models.Village{
		PlayerID:      player_id,
		VillageName:   fmt.Sprintf("Village %s", player_name),
		VillageSize:   100,
		VillageStatus: 0,
		VillageLocY:   int32(randInt(0, 100)),
		VillageLocX:   int32(randInt(0, 100)),
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

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	player_id, err := a.players.Insert(newPlayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "players.Insert"})
		return
	}

	village_id, err := a.createNewVillage(player_id, newPlayer.PlayerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "createNewVillage"})
		return
	}

	_, err = a.villageSetup.InsertWithIDCheck(village_id, player_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "function": "villageSetup.InsertWithIDCheck"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"player_id":  player_id,
		"village_id": village_id,
	})
}

// func postCalculateNewResources(c *gin.Context) {

// 	var numProductions models.Production

// 	if err := c.BindJSON(&numProductions); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	newResources := models.Resource{
// 		Food:   numProductions.HunterHut * 2,
// 		Wood:   numProductions.WoodcutterHut * 2,
// 		Stone:  numProductions.Quarry * 2,
// 		Copper: numProductions.CopperMine * 2,
// 		Water:  numProductions.Fountain * 2,
// 	}

// 	c.IndentedJSON(http.StatusOK, newResources)
// }
